package berth

func (b *Berth) Assign(shipID string) error {
	if b.busy {
		return errBusy
	}
	b.ShipID = shipID
	b.busy = true
	return nil
}

func (b *Berth) Clear() {
	b.busy = false
	b.ShipID = ""
}
