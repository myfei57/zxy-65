package berth

import "portpower/internal/meter"

func (b *Berth) Depart(m *meter.Meter) error {
	if !b.busy {
		return errNotBusy
	}
	b.Clear()
	m.Freeze()
	return nil
}
