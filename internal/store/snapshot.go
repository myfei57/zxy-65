package store

import "encoding/json"

func (s *Store) Dump() map[string]json.RawMessage {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make(map[string]json.RawMessage, len(s.data))
	for key, value := range s.data {
		out[key] = json.RawMessage(value)
	}
	return out
}

func (s *Store) Restore(entries map[string]json.RawMessage) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for key, value := range entries {
		if err := s.putLocked(key, []byte(value)); err != nil {
			return err
		}
	}
	return nil
}
