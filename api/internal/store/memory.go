// Package store implements the repositories. Today it is an in-memory store with
// seed data so you can run everything locally with zero AWS setup.
// Sprint 2 adds dynamodb.go implementing the same interfaces.
package store

import (
	"sort"
	"sync"

	"cakeshop/api/internal/cakes"
)

// Memory is a thread-safe in-memory repository.
type Memory struct {
	mu    sync.RWMutex
	cakes map[string]cakes.Cake
}

// NewMemory creates a store pre-filled with the given cakes.
func NewMemory(seed []cakes.Cake) *Memory {
	m := &Memory{cakes: make(map[string]cakes.Cake, len(seed))}
	for _, c := range seed {
		m.cakes[c.ID] = c
	}
	return m
}

// ListCakes returns all cakes sorted by name.
func (m *Memory) ListCakes() ([]cakes.Cake, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]cakes.Cake, 0, len(m.cakes))
	for _, c := range m.cakes {
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}

// GetCake returns one cake or cakes.ErrNotFound.
func (m *Memory) GetCake(id string) (cakes.Cake, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.cakes[id]
	if !ok {
		return cakes.Cake{}, cakes.ErrNotFound
	}
	return c, nil
}
