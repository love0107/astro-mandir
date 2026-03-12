package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/love0107/astro-mandir/internal/aztro"
)

type cache struct {
	data     map[string]*aztro.RashifalData
	cachedOn string
	mu       sync.Mutex
}

var rashiCache = &cache{
	data: make(map[string]*aztro.RashifalData),
}

type RashifalService struct{}

func NewRashifalService() *RashifalService {
	return &RashifalService{}
}

func (s *RashifalService) GetRashifal(ctx context.Context, rashi string) (*aztro.RashifalData, error) {
	_, ok := aztro.RashiMap[rashi]
	if !ok {
		return nil, fmt.Errorf("invalid rashi: %s", rashi)
	}

	today := time.Now().Format("2006-01-02")

	// Step 1 — check cache FIRST, then unlock immediately
	rashiCache.mu.Lock()
	if rashiCache.cachedOn == today {
		if cached, exists := rashiCache.data[rashi]; exists {
			rashiCache.mu.Unlock() // unlock before returning
			return cached, nil
		}
	} else {
		// New day — clear cache
		rashiCache.data = make(map[string]*aztro.RashifalData)
		rashiCache.cachedOn = today
	}
	rashiCache.mu.Unlock() // unlock BEFORE API call

	// Step 2 — API call happens outside lock
	// No other request is blocked during this
	data, err := aztro.FetchRashifal(rashi)
	if err != nil {
		return aztro.HardcodedRashifal(rashi), nil
	}

	// Step 3 — store in cache, lock again briefly
	rashiCache.mu.Lock()
	rashiCache.data[rashi] = data
	rashiCache.mu.Unlock()

	return data, nil
}
