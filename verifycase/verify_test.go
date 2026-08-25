package verifycase

import (
	"testing"

	"portpower/internal/disconn"
	"portpower/internal/meter"
	"portpower/internal/store"
)

func TestPpFinalReadingLost(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	binder := meter.NewBinder(st)
	if err := binder.Bind(meter.Binding{Berth: "A1", Ship: "s1"}); err != nil {
		t.Fatal(err)
	}
	breaker := disconn.NewBreaker("shore")
	if err := breaker.Reset(); err != nil {
		t.Fatal(err)
	}
	conn := disconn.Connection{Breaker: breaker, Binder: binder, Berth: "A1"}
	if err := disconn.BatchDisconnect([]disconn.Connection{conn}, map[string]float64{"A1": 150}); err != nil {
		t.Fatalf("batch disconnect failed: %v", err)
	}
	got, ok := binder.FinalValue("A1")
	if !ok || got != 150 {
		t.Fatalf("expected final value 150, got %v ok=%v", got, ok)
	}
}
