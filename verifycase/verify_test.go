package verifycase

import (
	"testing"

	"portpower/internal/berth"
	"portpower/internal/meter"
	"portpower/internal/store"
)

func TestPpMeterFreezeDurable(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	b := berth.NewBerth("A1")
	if err := b.Assign("s1"); err != nil {
		t.Fatal(err)
	}
	m := meter.NewMeter("m1", "A1")
	m.SetStore(st)
	if err := b.Depart(m); err != nil {
		t.Fatal(err)
	}
	reloaded := meter.NewMeter("m2", "A1")
	reloaded.SetStore(st)
	if !reloaded.Frozen() {
		t.Fatal("meter freeze must survive a reload")
	}
}
