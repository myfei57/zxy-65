package store

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type Store struct {
	mu   sync.Mutex
	dir  string
	data map[string][]byte
}

func Open(dir string) (*Store, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	s := &Store{dir: dir, data: map[string][]byte{}}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		raw, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			return nil, err
		}
		s.data[strings.TrimSuffix(entry.Name(), ".json")] = raw
	}
	return s, nil
}

func (s *Store) Put(key string, value []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.putLocked(key, value)
}

func (s *Store) putLocked(key string, value []byte) error {
	if err := os.WriteFile(filepath.Join(s.dir, key+".json"), value, 0o644); err != nil {
		return err
	}
	s.data[key] = value
	return nil
}

func (s *Store) Get(key string) ([]byte, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	raw, ok := s.data[key]
	return raw, ok
}

func (s *Store) Has(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.data[key]
	return ok
}

func (s *Store) Keys(prefix string) []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := []string{}
	for key := range s.data {
		if strings.HasPrefix(key, prefix) {
			out = append(out, key)
		}
	}
	sort.Strings(out)
	return out
}

func (s *Store) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := os.Remove(filepath.Join(s.dir, key+".json")); err != nil && !os.IsNotExist(err) {
		return err
	}
	delete(s.data, key)
	return nil
}
