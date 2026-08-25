package ship

import "sort"

type Fleet struct {
	ships map[string]*Ship
}

func NewFleet() *Fleet {
	return &Fleet{ships: map[string]*Ship{}}
}

func (f *Fleet) Add(id, name string) *Ship {
	if existing, ok := f.ships[id]; ok {
		return existing
	}
	created := NewShip(id, name)
	f.ships[id] = created
	return created
}

func (f *Fleet) Get(id string) (*Ship, bool) {
	s, ok := f.ships[id]
	return s, ok
}

func (f *Fleet) IDs() []string {
	out := make([]string, 0, len(f.ships))
	for id := range f.ships {
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}

func (f *Fleet) Count() int {
	return len(f.ships)
}

func (f *Fleet) Moored() []string {
	out := []string{}
	for _, s := range f.ships {
		if s.MooredAt() != "" {
			out = append(out, s.ID)
		}
	}
	sort.Strings(out)
	return out
}
