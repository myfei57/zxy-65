package meter

import "portpower/internal/store"

type Collector struct {
	store *store.Store
	seen  map[string]bool
}

func NewCollector(st *store.Store) *Collector {
	return &Collector{store: st, seen: map[string]bool{}}
}

func (c *Collector) Record(ts string, value float64) error {
	if c.seen[ts] {
		return nil
	}
	c.seen[ts] = true
	return c.store.AppendUsage(ts, value)
}
