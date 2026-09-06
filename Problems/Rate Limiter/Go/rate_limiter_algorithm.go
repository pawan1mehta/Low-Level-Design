package main

import "time"

type RateLimiterAlgorithm interface {
	Allow(time time.Time) Result
}
