package cable

type Monitor struct {
	filter *Filter
	peak   float64
	count  int
}

func NewMonitor(filter *Filter) *Monitor {
	return &Monitor{filter: filter}
}

func (m *Monitor) Observe(sample float64) {
	m.filter.Add(sample)
	if sample > m.peak {
		m.peak = sample
	}
	m.count++
}

func (m *Monitor) Peak() float64 {
	return m.peak
}

func (m *Monitor) Count() int {
	return m.count
}

func (m *Monitor) Average() float64 {
	return m.filter.Average()
}
