package plug

type State string

const (
	StateDisconnected State = "disconnected"
	StateEngaged      State = "engaged"
)

func (p *Plug) State() State {
	if p.engaged {
		return StateEngaged
	}
	return StateDisconnected
}
