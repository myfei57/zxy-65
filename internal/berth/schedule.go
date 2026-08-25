package berth

import (
	"encoding/json"
	"sort"

	"portpower/internal/store"
)

type Reservation struct {
	Berth  string `json:"berth"`
	Ship   string `json:"ship"`
	Window string `json:"window"`
}

type Scheduler struct {
	store *store.Store
}

func NewScheduler(st *store.Store) *Scheduler {
	return &Scheduler{store: st}
}

func reservationKey(berth string) string {
	return "berth.reservation." + berth
}

func (s *Scheduler) Book(r Reservation) error {
	raw, err := json.Marshal(r)
	if err != nil {
		return err
	}
	return s.store.Put(reservationKey(r.Berth), raw)
}

func (s *Scheduler) Get(berth string) (Reservation, bool) {
	raw, ok := s.store.Get(reservationKey(berth))
	if !ok {
		return Reservation{}, false
	}
	var r Reservation
	if err := json.Unmarshal(raw, &r); err != nil {
		return Reservation{}, false
	}
	return r, true
}

func (s *Scheduler) Cancel(berth string) error {
	return s.store.Delete(reservationKey(berth))
}

func (s *Scheduler) List() []Reservation {
	out := []Reservation{}
	for _, key := range s.store.Keys("berth.reservation.") {
		code := key[len("berth.reservation."):]
		if r, ok := s.Get(code); ok {
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Berth < out[j].Berth })
	return out
}
