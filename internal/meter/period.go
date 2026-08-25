package meter

import "portpower/internal/store"

type Period struct {
	Start string  `json:"start"`
	End   string  `json:"end"`
	Total float64 `json:"total"`
	Count int     `json:"count"`
}

func Aggregate(st *store.Store, start, end string) Period {
	result := Period{Start: start, End: end}
	for _, rec := range st.UsageRecords() {
		if rec.Timestamp >= start && rec.Timestamp <= end {
			result.Total += rec.Value
			result.Count++
		}
	}
	return result
}
