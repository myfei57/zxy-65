package berth

import "sort"

type Registry struct {
	berths map[string]*Berth
}

func NewRegistry() *Registry {
	return &Registry{berths: map[string]*Berth{}}
}

func (r *Registry) Add(code string) *Berth {
	if existing, ok := r.berths[code]; ok {
		return existing
	}
	created := NewBerth(code)
	r.berths[code] = created
	return created
}

func (r *Registry) Get(code string) (*Berth, bool) {
	b, ok := r.berths[code]
	return b, ok
}

func (r *Registry) Codes() []string {
	out := make([]string, 0, len(r.berths))
	for code := range r.berths {
		out = append(out, code)
	}
	sort.Strings(out)
	return out
}

func (r *Registry) OccupiedCount() int {
	count := 0
	for _, b := range r.berths {
		if b.Occupied() {
			count++
		}
	}
	return count
}

func (r *Registry) Clear(code string) bool {
	b, ok := r.berths[code]
	if !ok {
		return false
	}
	b.Clear()
	return true
}
