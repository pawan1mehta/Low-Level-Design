package main

import "time"

type Algorithm int

const (
	TokenBucket Algorithm = iota
	SlidingWindow
)

type AlgoConfig interface {
	Algorithm() Algorithm
}

type BucketAlgoConfig struct {
	capacity        int
	refillPerSecond float64
}

func (BucketAlgoConfig) Algorithm() Algorithm {
	return TokenBucket
}

func (b BucketAlgoConfig) Capacity() int {
	return b.capacity
}

func (b BucketAlgoConfig) RefillPerSecond() float64 {
	return b.refillPerSecond
}

type SlidingWindowAlgoConfig struct {
	limit  int
	window time.Duration
}

func (SlidingWindowAlgoConfig) Algorithm() Algorithm {
	return SlidingWindow
}

func (b SlidingWindowAlgoConfig) Limit() int {
	return b.limit
}

func (b SlidingWindowAlgoConfig) Window() time.Duration {
	return b.window
}

type EndpointConfiguration struct {
	Algorithm Algorithm
	Config    AlgoConfig
}
