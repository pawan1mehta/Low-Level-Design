package main

type AlgorithmFactory struct {
}

func (af AlgorithmFactory) Create(config EndpointConfiguration) (RateLimiterAlgorithm, error) {
	switch config.Algorithm {
	case TokenBucket:
		config, ok := config.Config.(BucketAlgoConfig)
		if !ok {
			return nil, ErrConfigMismatch
		}
		return NewTokenBucketAlgorithm(config.Capacity(), config.RefillPerSecond()), nil

	case SlidingWindow:
		config, ok := config.Config.(SlidingWindowAlgoConfig)
		if !ok {
			return nil, ErrConfigMismatch
		}
		return NewSlidingWindowAlgorithm(config.Limit(), config.Window()), nil

	default:
		return nil, ErrUnknowConfigType
	}
}
