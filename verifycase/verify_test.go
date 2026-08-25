package verifycase

import (
	"testing"

	"portpower/internal/disconn"
	"portpower/internal/ship"
	"portpower/internal/store"
)

func TestPpShipProfileFresh(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	reg := ship.NewRegistry(st)
	if err := reg.Save(ship.Profile{ShipID: "s1", MaxLoad: 100, PhaseSeq: "ABC", Frequency: 50}); err != nil {
		t.Fatal(err)
	}
	enforcer := disconn.NewLimitEnforcer(reg)
	if !enforcer.Allow("s1", 80) {
		t.Fatal("expected initial allowance")
	}
	if err := reg.Retrofit("s1", 50); err != nil {
		t.Fatal(err)
	}
	if !enforcer.Allow("s1", 120) {
		t.Fatal("expected allowance after retrofit")
	}
}
