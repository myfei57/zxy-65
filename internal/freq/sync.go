package freq

import (
	"encoding/json"

	"portpower/internal/store"
)

func freqKey() string {
	return "freq.current"
}

func SetCurrent(st *store.Store, value float64) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return st.Put(freqKey(), raw)
}

func (s *Syncer) Refresh(value float64) error {
	if err := SetCurrent(s.store, value); err != nil {
		return err
	}
	s.value = value
	return nil
}

func (s *Syncer) Current() float64 {
	raw, ok := s.store.Get(freqKey())
	if !ok {
		return s.value
	}
	var value float64
	if err := json.Unmarshal(raw, &value); err != nil {
		return s.value
	}
	return value
}
