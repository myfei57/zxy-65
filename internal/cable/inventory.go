package cable

import "sort"

type Inventory struct {
	cables map[string]*Cable
}

func NewInventory() *Inventory {
	return &Inventory{cables: map[string]*Cable{}}
}

func (i *Inventory) Add(id string, rated float64) *Cable {
	if existing, ok := i.cables[id]; ok {
		return existing
	}
	created := NewCable(id, rated)
	i.cables[id] = created
	return created
}

func (i *Inventory) Get(id string) (*Cable, bool) {
	c, ok := i.cables[id]
	return c, ok
}

func (i *Inventory) IDs() []string {
	out := make([]string, 0, len(i.cables))
	for id := range i.cables {
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}

func (i *Inventory) Count() int {
	return len(i.cables)
}

func (i *Inventory) Select(minRated float64) (string, bool) {
	for _, id := range i.IDs() {
		c, _ := i.Get(id)
		if c.RatedTemperature() >= minRated {
			return id, true
		}
	}
	return "", false
}
