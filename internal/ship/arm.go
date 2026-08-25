package ship

type BreakerArm struct {
	armed bool
}

func NewBreakerArm() *BreakerArm {
	return &BreakerArm{}
}

func (a *BreakerArm) Arm() {
	a.armed = true
}

func (a *BreakerArm) Armed() bool {
	return a.armed
}
