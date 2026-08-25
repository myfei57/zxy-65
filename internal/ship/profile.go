package ship

import (
	"encoding/json"

	"portpower/internal/store"
)

type Profile struct {
	ShipID    string  `json:"ship"`
	MaxLoad   float64 `json:"max_load"`
	PhaseSeq  string  `json:"phase_seq"`
	Frequency float64 `json:"frequency"`
}

type Registry struct {
	store *store.Store
}

func NewRegistry(st *store.Store) *Registry {
	return &Registry{store: st}
}

func profileKey(ship string) string {
	return "ship.profile." + ship
}

func (r *Registry) Save(p Profile) error {
	raw, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return r.store.Put(profileKey(p.ShipID), raw)
}

func (r *Registry) Load(ship string) (Profile, bool) {
	raw, ok := r.store.Get(profileKey(ship))
	if !ok {
		return Profile{}, false
	}
	var p Profile
	if err := json.Unmarshal(raw, &p); err != nil {
		return Profile{}, false
	}
	return p, true
}
