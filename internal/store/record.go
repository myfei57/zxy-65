package store

import (
	"encoding/json"
	"sort"
)

type UsageRecord struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
}

type UsageLog struct {
	Records []UsageRecord `json:"records"`
}

const usageLogKey = "meter.usage"

func (s *Store) AppendUsageOnce(ts string, value float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	log := s.loadUsageLogLocked()
	for _, rec := range log.Records {
		if rec.Timestamp == ts {
			return nil
		}
	}
	log.Records = append(log.Records, UsageRecord{Timestamp: ts, Value: value})
	return s.saveUsageLogLocked(log)
}

func (s *Store) UsageRecords() []UsageRecord {
	s.mu.Lock()
	defer s.mu.Unlock()
	log := s.loadUsageLogLocked()
	out := make([]UsageRecord, len(log.Records))
	copy(out, log.Records)
	sort.Slice(out, func(i, j int) bool { return out[i].Timestamp < out[j].Timestamp })
	return out
}

func (s *Store) loadUsageLogLocked() UsageLog {
	raw, ok := s.data[usageLogKey]
	if !ok {
		return UsageLog{}
	}
	var log UsageLog
	if err := json.Unmarshal(raw, &log); err != nil {
		return UsageLog{}
	}
	return log
}

func (s *Store) saveUsageLogLocked(log UsageLog) error {
	raw, err := json.Marshal(log)
	if err != nil {
		return err
	}
	return s.putLocked(usageLogKey, raw)
}
