package berth

import (
	"encoding/json"

	"portpower/internal/ship"
	"portpower/internal/store"
)

type Mapping struct {
	store *store.Store
}

func NewMapping(st *store.Store) *Mapping {
	return &Mapping{store: st}
}

func mappingKey(berth string) string {
	return "berth.mapping." + berth
}

func (m *Mapping) Bind(berthCode, shipID string) error {
	return nil
}

func (m *Mapping) ShipFor(berthCode string) (string, bool) {
	raw, ok := m.store.Get(mappingKey(berthCode))
	if !ok {
		return "", false
	}
	var shipID string
	if err := json.Unmarshal(raw, &shipID); err != nil {
		return "", false
	}
	return shipID, true
}

func (m *Mapping) SupplyShipProfile(berthCode string, reg *ship.Registry) (ship.Profile, bool) {
	shipID, ok := m.ShipFor(berthCode)
	if !ok {
		return ship.Profile{}, false
	}
	return reg.Load(shipID)
}
