package quota

func (m *Manager) Remaining(ship string) float64 {
	q, ok := m.Get(ship)
	if !ok {
		return 0
	}
	return q.Limit - q.Used
}
