package safemap

import "sync"

type SafeMap[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex
}

func (s *SafeMap[K, V]) Put(key K, val V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = val
	//
}

func (s *SafeMap[K, V]) Get(key K) (V, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.m[key]
	return val, ok
}

func (s *SafeMap[K, V]) LoadOrStore(key K, newVal V) (val V, loaded bool) {
	s.mu.RLock()
	v, ok := s.m[key]
	s.mu.RUnlock()
	if ok {
		return v, ok
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// double-check
	if v, ok = s.m[key]; ok {
		return v, ok
	}

	s.m[key] = newVal

	return newVal, false
}
