package quota

func (m *Manager) Reserve(ship string, demand float64) bool {
	q, ok := m.Get(ship)
	if !ok {
		return false
	}
	if demand+q.Used > q.Limit {
		return false
	}
	q.Used += demand
	return m.Set(q) == nil
}

func (m *Manager) Release(ship string, amount float64) bool {
	q, ok := m.Get(ship)
	if !ok {
		return false
	}
	if amount > q.Used {
		q.Used = 0
	} else {
		q.Used -= amount
	}
	return m.Set(q) == nil
}
