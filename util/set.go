package util

import "sync"

// NewSet ...
func NewSet() *Set {
	m := map[string]struct{}{}
	s := &Set{}
	s.m = m
	return s
}

// Set ...
type Set struct {
	m   map[string]struct{}
	mux sync.Mutex
}

// List ...
func (s *Set) List() []string {
	s.mux.Lock()
	defer s.mux.Unlock()
	r := []string{}
	for k := range s.m {
		r = append(r, k)
	}
	return r
}

// Add ...
func (s *Set) Add(v string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m[v] = struct{}{}
}
