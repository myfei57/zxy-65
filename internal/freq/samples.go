package freq

import (
	"encoding/json"
)

const samplesKey = "freq.samples"

func (s *Syncer) RecordSample(value float64) error {
	samples := s.Samples()
	samples = append(samples, value)
	raw, err := json.Marshal(samples)
	if err != nil {
		return err
	}
	return s.store.Put(samplesKey, raw)
}

func (s *Syncer) Samples() []float64 {
	raw, ok := s.store.Get(samplesKey)
	if !ok {
		return nil
	}
	var samples []float64
	if err := json.Unmarshal(raw, &samples); err != nil {
		return nil
	}
	return samples
}

func (s *Syncer) Within(target, tolerance float64) bool {
	return InBand(s.Current(), target, tolerance)
}
