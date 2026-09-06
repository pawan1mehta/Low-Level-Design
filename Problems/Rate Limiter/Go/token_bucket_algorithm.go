package main

import (
	"math"
	"sync"
	"time"
)

type TokenBucketAlgorithm struct {
	capacity        int
	tokens          float64
	refillPerSecond float64
	lastRefill      time.Time
	mu              sync.Mutex
}

func NewTokenBucketAlgorithm(capacity int, refillPerSecond float64) *TokenBucketAlgorithm {
	return &TokenBucketAlgorithm{
		capacity:        capacity,
		tokens:          float64(capacity),
		refillPerSecond: refillPerSecond,
		lastRefill:      time.Now(),
		mu:              sync.Mutex{},
	}
}

func (t *TokenBucketAlgorithm) Allow(now time.Time) Result {
	t.mu.Lock()
	defer t.mu.Unlock()

	elapsed := now.Sub(t.lastRefill).Seconds()
	if elapsed < 0 {
		elapsed = 0
	}

	t.tokens = math.Min(float64(t.capacity), t.tokens+elapsed*t.refillPerSecond)
	t.lastRefill = now

	if t.tokens >= 1 {
		t.tokens -= 1
		return Result{
			Allowed:    true,
			Remaining:  int64(t.tokens),
			RetryAfter: 0,
		}
	}

	retryAfter := 0
	if t.refillPerSecond > 0 {
		retryAfter = int(math.Ceil((1 - t.tokens) / t.refillPerSecond))
	}

	return Result{
		Allowed:    false,
		Remaining:  0,
		RetryAfter: retryAfter,
	}
}
