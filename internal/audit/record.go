package audit

func (l *Ledger) Count() int {
	return len(l.Entries())
}

func (l *Ledger) Latest() (Entry, bool) {
	entries := l.Entries()
	if len(entries) == 0 {
		return Entry{}, false
	}
	return entries[len(entries)-1], true
}
