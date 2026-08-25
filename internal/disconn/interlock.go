package disconn

import (
	"portpower/internal/phase"
	"portpower/internal/ship"
)

func (b *Breaker) InterlockReady(arm *ship.BreakerArm, verifier *phase.Verifier) bool {
	if !arm.Armed() {
		return false
	}
	if err := verifier.Verify(); err != nil {
		return false
	}
	return true
}
