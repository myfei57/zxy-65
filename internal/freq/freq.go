package freq

import "portpower/internal/store"

type Syncer struct {
	store *store.Store
	value float64
}

func NewSyncer(st *store.Store, initial float64) *Syncer {
	return &Syncer{store: st, value: initial}
}
