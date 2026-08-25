package store

import "strings"

func (s *Store) Values(prefix string) map[string][]byte {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := map[string][]byte{}
	for key, value := range s.data {
		if strings.HasPrefix(key, prefix) {
			out[key] = value
		}
	}
	return out
}

func (s *Store) Count(prefix string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	count := 0
	for key := range s.data {
		if strings.HasPrefix(key, prefix) {
			count++
		}
	}
	return count
}
