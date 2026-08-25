package plug

import "errors"

var errNotEngaged = errors.New("plug is not engaged")

type Plug struct {
	ID        string
	engaged   bool
	boundShip string
}

func NewPlug(id string) *Plug {
	return &Plug{ID: id}
}

func (p *Plug) Engaged() bool {
	return p.engaged
}

func (p *Plug) BoundShip() string {
	return p.boundShip
}

func (p *Plug) markEngaged() {
	p.engaged = true
}
