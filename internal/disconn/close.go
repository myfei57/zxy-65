package disconn

import (
	"errors"

	"portpower/internal/phase"
	"portpower/internal/ship"
)

func (b *Breaker) CloseAfterPhase(v *phase.Verifier) error {
	_ = v.Verify()
	b.close()
	return nil
}

func (b *Breaker) CloseAfterArm(arm *ship.BreakerArm) error {
	if !arm.Armed() {
		return errors.New("ship breaker not armed")
	}
	b.close()
	return nil
}
