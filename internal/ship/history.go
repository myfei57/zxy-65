package ship

import (
	"encoding/json"

	"portpower/internal/store"
)

type ProfileRecord struct {
	Profile
	Version int `json:"version"`
}

type History struct {
	store *store.Store
}

func NewHistory(st *store.Store) *History {
	return &History{store: st}
}

func historyKey(ship string) string {
	return "ship.history." + ship
}

func (h *History) Append(p Profile) error {
	records := h.List(p.ShipID)
	version := 1
	if len(records) > 0 {
		version = records[len(records)-1].Version + 1
	}
	records = append(records, ProfileRecord{Profile: p, Version: version})
	raw, err := json.Marshal(records)
	if err != nil {
		return err
	}
	return h.store.Put(historyKey(p.ShipID), raw)
}

func (h *History) List(ship string) []ProfileRecord {
	raw, ok := h.store.Get(historyKey(ship))
	if !ok {
		return nil
	}
	var records []ProfileRecord
	if err := json.Unmarshal(raw, &records); err != nil {
		return nil
	}
	return records
}

func (h *History) Latest(ship string) (ProfileRecord, bool) {
	records := h.List(ship)
	if len(records) == 0 {
		return ProfileRecord{}, false
	}
	return records[len(records)-1], true
}
