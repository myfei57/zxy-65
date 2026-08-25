package cable

type Sensor struct {
	value float64
}

func NewSensor(initial float64) *Sensor {
	return &Sensor{value: initial}
}

func (s *Sensor) Read() float64 {
	return s.value
}

func (s *Sensor) Set(value float64) {
	s.value = value
}
