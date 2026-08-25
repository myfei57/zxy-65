package quota

func (m *Manager) Allow(ship string, demand float64) bool {
	q, ok := m.Get(ship)
	if !ok {
		return false
	}
	return demand <= q.Limit
}
