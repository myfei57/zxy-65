package meter

import "portpower/internal/store"

type Collector struct {
	store *store.Store
}

func NewCollector(st *store.Store) *Collector {
	return &Collector{store: st}
}

func (c *Collector) Record(ts string, value float64) error {
	return c.store.AppendUsageOnce(ts, value)
}
