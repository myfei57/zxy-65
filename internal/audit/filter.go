package audit

func (l *Ledger) ByEvent(event string) []Entry {
	out := []Entry{}
	for _, entry := range l.Entries() {
		if entry.Event == event {
			out = append(out, entry)
		}
	}
	return out
}

func (l *Ledger) Events() []string {
	seen := map[string]bool{}
	out := []string{}
	for _, entry := range l.Entries() {
		if !seen[entry.Event] {
			seen[entry.Event] = true
			out = append(out, entry.Event)
		}
	}
	return out
}
