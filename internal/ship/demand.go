package ship

func (p Profile) Valid() bool {
	return p.ShipID != "" && p.MaxLoad > 0 && p.Frequency > 0
}

func (p Profile) PhaseValid() bool {
	return p.PhaseSeq == "ABC" || p.PhaseSeq == "CBA"
}

func (r *Registry) Demand(ship string) float64 {
	p, ok := r.Load(ship)
	if !ok {
		return 0
	}
	return p.MaxLoad
}

func (r *Registry) Headroom(ship string) float64 {
	return r.Demand(ship) * 1.1
}
