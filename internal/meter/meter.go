package meter

import "portpower/internal/store"

type Meter struct {
	ID     string
	Berth  string
	frozen bool
	store  *store.Store
}

func NewMeter(id, berth string) *Meter {
	return &Meter{ID: id, Berth: berth}
}

func (m *Meter) SetStore(st *store.Store) {
	m.store = st
}

func (m *Meter) Frozen() bool {
	if m.store != nil {
		if _, ok := m.store.Get("meter.frozen." + m.Berth); ok {
			return true
		}
	}
	return m.frozen
}
