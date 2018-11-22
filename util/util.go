package util

import "sync"

// Separate ...
func Separate(chars []string, idx int) (string, []string) {
	if idx < 0 {
		return "", []string{}
	}
	if chars == nil {
		return "", []string{}
	}
	if len(chars) == 0 {
		return "", []string{}
	}
	if len(chars) <= idx {
		return "", []string{}
	}
	var target string
	var remaining []string
	for i, c := range chars {
		if i == idx {
			target = c
		} else {
			remaining = append(remaining, c)
		}
	}
	return target, remaining
}

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
	// ignore worning
	for k, _ := range s.m {
		// fmt.Println(k, v)
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
