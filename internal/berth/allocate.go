package berth

func (r *Registry) Allocate(shipID string) (string, bool) {
	for _, code := range r.Codes() {
		b, _ := r.Get(code)
		if !b.Occupied() {
			if err := b.Assign(shipID); err != nil {
				continue
			}
			return code, true
		}
	}
	return "", false
}

func (r *Registry) Vacate(code string) bool {
	return r.Clear(code)
}

func (r *Registry) ShipAt(code string) (string, bool) {
	b, ok := r.Get(code)
	if !ok {
		return "", false
	}
	if !b.Occupied() {
		return "", false
	}
	return b.ShipID, true
}
