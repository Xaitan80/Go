package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// cacheEntry represents a cached Pokémon entry.
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// MockCache implements the Cache interface for testing purposes
type MockCache struct {
	Data     map[string]*cacheEntry
	mu       sync.RWMutex
	interval time.Duration
}

func (m *MockCache) Get(name string) ([]byte, bool) {
	m.mu.RLock()
	entry, ok := m.Data[name]
	m.mu.RUnlock()

	if !ok {
		fmt.Printf("Pokémon '%s' not found in cache. Returning default.\n", name)
		return nil, false
	}

	return entry.val, true
}

func (m *MockCache) Set(name string, val []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()
	entry := &cacheEntry{createdAt: time.Now()}
	entry.val = val
	// entry.name = name removed
	m.Data[name] = entry
}

func TestCacheFunctionality(t *testing.T) {
	// Create a mock cache instance
	mockCache := &MockCache{
		Data:     make(map[string]*cacheEntry),
		mu:       sync.RWMutex{},
		interval: 1 * time.Second,
	}

	// Example data to be stored in the cache
	pokemonName := "Pikachu"
	data := []byte("This is the data for " + pokemonName)

	// Store data in the cache
	mockCache.Set(pokemonName, data)

	// Retrieve data from the cache
	retrievedData, found := mockCache.Get(pokemonName)

	// Assertions
	if !found {
		t.Errorf("Cache miss for %s", pokemonName)
		return
	}

	if retrievedData == nil {
		t.Errorf("Retrieved data for %s is nil", pokemonName)
		return
	}

	if string(retrievedData) != "This is the data for "+pokemonName {
		t.Errorf("Data mismatch for %s: expected '%s', got '%s'", pokemonName, "This is the data for "+pokemonName, string(retrievedData))
	}

	fmt.Printf("Cache test passed!\n")
}
