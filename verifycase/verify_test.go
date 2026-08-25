package verifycase

import (
	"testing"

	"portpower/internal/meter"
	"portpower/internal/store"
)

func TestPpMeterDedupDurable(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	a := meter.NewCollector(st)
	b := meter.NewCollector(st)
	if err := a.Record("T1", 100); err != nil {
		t.Fatal(err)
	}
	if err := b.Record("T1", 100); err != nil {
		t.Fatal(err)
	}
	records := st.UsageRecords()
	if len(records) != 1 {
		t.Fatalf("expected 1 usage record, got %d", len(records))
	}
}
