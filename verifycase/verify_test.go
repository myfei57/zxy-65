package verifycase

import (
	"testing"

	"portpower/internal/disconn"
	"portpower/internal/freq"
	"portpower/internal/store"
)

func TestPpFreqSyncFresh(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	sync := freq.NewSyncer(st, 50)
	if err := freq.SetCurrent(st, 60); err != nil {
		t.Fatal(err)
	}
	breaker := disconn.NewBreaker("b1")
	if err := breaker.ParallelClose(sync, 60, 0.1); err != nil {
		t.Fatalf("expected parallel close to use fresh frequency: %v", err)
	}
	if !breaker.Closed() {
		t.Fatal("breaker must close when frequency is in sync")
	}
}
