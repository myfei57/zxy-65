package verifycase

import (
	"testing"

	"portpower/internal/meter"
	"portpower/internal/plug"
	"portpower/internal/store"
)

func TestPpPlugBindingStale(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	binder1 := meter.NewBinder(st)
	if err := binder1.Bind(meter.Binding{Berth: "A1", Ship: "old"}); err != nil {
		t.Fatal(err)
	}
	binder2 := meter.NewBinder(st)
	if err := binder2.Bind(meter.Binding{Berth: "A1", Ship: "new"}); err != nil {
		t.Fatal(err)
	}
	connector := plug.NewPlug("p1")
	if err := connector.Engage("A1", binder1); err != nil {
		t.Fatal(err)
	}
	if connector.BoundShip() != "new" {
		t.Fatalf("expected bound ship new, got %s", connector.BoundShip())
	}
}
