package verifycase

import (
	"testing"

	"portpower/internal/berth"
	"portpower/internal/ship"
	"portpower/internal/store"
)

func TestPpBerthMappingFresh(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	reg := ship.NewRegistry(st)
	if err := reg.Save(ship.Profile{ShipID: "s1", MaxLoad: 100, PhaseSeq: "ABC", Frequency: 50}); err != nil {
		t.Fatal(err)
	}
	if err := reg.Save(ship.Profile{ShipID: "s2", MaxLoad: 200, PhaseSeq: "ABC", Frequency: 50}); err != nil {
		t.Fatal(err)
	}
	mapping := berth.NewMapping(st)
	if err := mapping.Bind("A1", "s1"); err != nil {
		t.Fatal(err)
	}
	if err := mapping.Bind("A1", "s2"); err != nil {
		t.Fatal(err)
	}
	profile, ok := mapping.SupplyShipProfile("A1", reg)
	if !ok || profile.ShipID != "s2" {
		t.Fatalf("expected supply profile s2, got %+v ok=%v", profile, ok)
	}
}
