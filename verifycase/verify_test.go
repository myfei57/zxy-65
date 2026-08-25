package verifycase

import (
	"testing"

	"portpower/internal/disconn"
	"portpower/internal/phase"
)

func TestPpPhaseVerifySwallowed(t *testing.T) {
	verifier := phase.NewVerifier(phase.SequenceReverse)
	breaker := disconn.NewBreaker("b1")
	err := breaker.CloseAfterPhase(verifier)
	if err == nil {
		t.Fatal("expected reversed phase to be rejected")
	}
	if breaker.Closed() {
		t.Fatal("breaker must stay open on reversed phase")
	}
}
