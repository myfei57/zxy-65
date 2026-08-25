package audit

type Summary struct {
	Total   int            `json:"total"`
	ByEvent map[string]int `json:"by_event"`
}

func (l *Ledger) Summary() Summary {
	entries := l.Entries()
	byEvent := map[string]int{}
	for _, entry := range entries {
		byEvent[entry.Event]++
	}
	return Summary{Total: len(entries), ByEvent: byEvent}
}
