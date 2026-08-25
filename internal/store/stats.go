package store

func (s *Store) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.data)
}

func (s *Store) TotalBytes() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	total := 0
	for _, value := range s.data {
		total += len(value)
	}
	return total
}
