package berth

import "errors"

var errBusy = errors.New("berth already occupied")
var errNotBusy = errors.New("berth is not occupied")

type Berth struct {
	Code   string
	ShipID string
	busy   bool
}

func NewBerth(code string) *Berth {
	return &Berth{Code: code}
}

func (b *Berth) Occupied() bool {
	return b.busy
}
