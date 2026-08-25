package store

import "encoding/json"

func (s *Store) ExportJSON() ([]byte, error) {
	return json.Marshal(s.Dump())
}

func (s *Store) ImportJSON(raw []byte) error {
	var entries map[string]json.RawMessage
	if err := json.Unmarshal(raw, &entries); err != nil {
		return err
	}
	return s.Restore(entries)
}
