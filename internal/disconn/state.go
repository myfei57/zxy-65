package disconn

type State string

const (
	StateOpen    State = "open"
	StateClosed  State = "closed"
	StateTripped State = "tripped"
)

func (b *Breaker) State() State {
	if b.tripped {
		return StateTripped
	}
	if b.closed {
		return StateClosed
	}
	return StateOpen
}

func (b *Breaker) Describe() string {
	return string(b.State())
}
