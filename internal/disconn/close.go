package disconn

import (
	"portpower/internal/phase"
	"portpower/internal/ship"
)

func (b *Breaker) CloseAfterPhase(v *phase.Verifier) error {
	if err := v.Verify(); err != nil {
		return err
	}
	b.close()
	return nil
}

func (b *Breaker) CloseAfterArm(arm *ship.BreakerArm) error {
	if !arm.Armed() {
		return errShipNotArmed
	}
	b.close()
	return nil
}
