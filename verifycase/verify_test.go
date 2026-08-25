package verifycase

import (
	"testing"

	"portpower/internal/disconn"
	"portpower/internal/ship"
)

func TestPpArmCheckSwallowed(t *testing.T) {
	arm := ship.NewBreakerArm()
	breaker := disconn.NewBreaker("shore")
	err := breaker.CloseAfterArm(arm)
	if err == nil {
		t.Fatal("expected unarmed ship breaker to be rejected")
	}
	if breaker.Closed() {
		t.Fatal("shore breaker must stay open until ship breaker arms")
	}
}
