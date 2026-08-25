package berth

type Status struct {
	Code     string `json:"code"`
	Occupied bool   `json:"occupied"`
	ShipID   string `json:"ship_id,omitempty"`
}

func (r *Registry) Status(code string) (Status, bool) {
	b, ok := r.Get(code)
	if !ok {
		return Status{}, false
	}
	item := Status{Code: code, Occupied: b.Occupied()}
	if b.Occupied() {
		item.ShipID = b.ShipID
	}
	return item, true
}

func (r *Registry) StatusAll() []Status {
	out := []Status{}
	for _, code := range r.Codes() {
		if item, ok := r.Status(code); ok {
			out = append(out, item)
		}
	}
	return out
}
