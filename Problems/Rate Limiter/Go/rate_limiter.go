package main

import (
	"sync"
	"time"
)

type RateLimiter struct {
	configurationStore *ConfigurationStore
	factory            AlgorithmFactory
	states             map[Key]RateLimiterAlgorithm
	mu                 sync.Mutex
}

func NewRateLimiter(store *ConfigurationStore, factory AlgorithmFactory) *RateLimiter {
	return &RateLimiter{
		configurationStore: store,
		factory:            factory,
		states:             map[Key]RateLimiterAlgorithm{},
		mu:                 sync.Mutex{},
	}
}

func (r *RateLimiter) RegisterEndpoints(endpoint string, config EndpointConfiguration) error {
	return r.configurationStore.AddEndpoint(endpoint, config)
}

func (r *RateLimiter) Check(clientID string, endpoint string) (Result, error) {
	endpointConfig, _ := r.configurationStore.GetConfig(endpoint)
	key := NewKey(clientID, endpoint)
	r.mu.Lock()
	_, ok := r.states[key]
	if !ok {
		algo, err := r.factory.Create(endpointConfig)
		if err != nil {
			return Result{}, err
		}
		r.states[key] = algo
	}
	r.mu.Unlock()
	return r.states[key].Allow(time.Now()), nil
}
