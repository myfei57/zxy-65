package quota

import (
	"encoding/json"

	"portpower/internal/store"
)

type Quota struct {
	ShipID string  `json:"ship"`
	Limit  float64 `json:"limit"`
	Used   float64 `json:"used"`
}

type Manager struct {
	store *store.Store
}

func NewManager(st *store.Store) *Manager {
	return &Manager{store: st}
}

func quotaKey(ship string) string {
	return "quota." + ship
}

func (m *Manager) Set(q Quota) error {
	raw, err := json.Marshal(q)
	if err != nil {
		return err
	}
	return m.store.Put(quotaKey(q.ShipID), raw)
}

func (m *Manager) Get(ship string) (Quota, bool) {
	raw, ok := m.store.Get(quotaKey(ship))
	if !ok {
		return Quota{}, false
	}
	var q Quota
	if err := json.Unmarshal(raw, &q); err != nil {
		return Quota{}, false
	}
	return q, true
}
