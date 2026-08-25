package meter

import "portpower/internal/store"

func (m *Meter) CanRecord() bool {
	return !m.frozen
}

func TotalUsage(st *store.Store) float64 {
	total := 0.0
	for _, rec := range st.UsageRecords() {
		total += rec.Value
	}
	return total
}
