package meter

import "portpower/internal/store"

type Tariff struct {
	Rate float64
}

func NewTariff(rate float64) *Tariff {
	return &Tariff{Rate: rate}
}

func (t *Tariff) Charge(total float64) float64 {
	return total * t.Rate
}

func Bill(st *store.Store, rate float64) float64 {
	return TotalUsage(st) * rate
}
