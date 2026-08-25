package cable

type Cable struct {
	ID     string
	ratedC float64
}

func NewCable(id string, ratedC float64) *Cable {
	return &Cable{ID: id, ratedC: ratedC}
}

func (c *Cable) RatedTemperature() float64 {
	return c.ratedC
}
