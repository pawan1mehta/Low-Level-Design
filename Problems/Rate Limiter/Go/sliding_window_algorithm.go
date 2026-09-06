package main

import (
	"sync"
	"time"
)

type SlidingWindowAlgorithm struct {
	limit      int
	window     time.Duration
	timestamps []time.Time
	mu         sync.Mutex
}

func NewSlidingWindowAlgorithm(limit int, window time.Duration) *SlidingWindowAlgorithm {
	return &SlidingWindowAlgorithm{
		limit:      limit,
		window:     window,
		timestamps: []time.Time{},
		mu:         sync.Mutex{},
	}
}

func (t *SlidingWindowAlgorithm) Allow(now time.Time) Result {
	// TODO: implement me
	return Result{}
}
