package plug

import (
	"errors"

	"portpower/internal/meter"
)

var errNoBinding = errors.New("no meter binding for berth")

func (p *Plug) Engage(berth string, binder *meter.Binder) error {
	bind, ok := binder.Load(berth)
	if !ok {
		return errNoBinding
	}
	p.boundShip = bind.Ship
	p.markEngaged()
	return nil
}
