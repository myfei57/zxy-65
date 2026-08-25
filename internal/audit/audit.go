package audit

import (
	"encoding/json"
	"sort"
	"time"

	"portpower/internal/store"
)

type Entry struct {
	At     string `json:"at"`
	Event  string `json:"event"`
	Detail string `json:"detail"`
}

type Ledger struct {
	store *store.Store
}

const auditKey = "audit.log"

func NewLedger(st *store.Store) *Ledger {
	return &Ledger{store: st}
}

func (l *Ledger) Append(event, detail string) error {
	entries := l.Entries()
	entries = append(entries, Entry{At: time.Now().UTC().Format(time.RFC3339), Event: event, Detail: detail})
	raw, err := json.Marshal(entries)
	if err != nil {
		return err
	}
	return l.store.Put(auditKey, raw)
}

func (l *Ledger) Entries() []Entry {
	raw, ok := l.store.Get(auditKey)
	if !ok {
		return nil
	}
	var entries []Entry
	if err := json.Unmarshal(raw, &entries); err != nil {
		return nil
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].At < entries[j].At })
	return entries
}
