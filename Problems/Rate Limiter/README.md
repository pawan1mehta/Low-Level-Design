# Rate Limiter

## Clarifying Questions

- How many algorithms we'll support? [Ans: We'll support multiple algorithms]
- How many endpoints are we going to support?
- What infromation do we recieve? ClientID & endpoint or someting else? [Ans: Each request will contain ClientID & endpoint]
- What should we return when checking the reuest? [Ans: how many requests remain in their quota, and if denied when they can try]
- What should we do if we receive request for which we don't have any configuration? [Ans: fallback to default configurations]
- Will the rate limiter be distributed rate limiter across multiple server or in-memory? [Ans: In-memory]
- Is the configuration dynamic, or is it loaded at the startup?

## Requirements

- System receives requests with clientID & endpoint
- System support Multiple algorithms.
  - Each algorithm will have different configurations
- Each endpoing as configuration specifying:
  - Algorithm to use (ex: "Token Bucket", "Sliding Window", etc)
  - Algorithm specific paramters (ex: capacity, refillRatePerSecond for Token...)
- System enforces rate limits by checking the clientId and endpoint
- If endpoint has no configuration set,then use the default configuration
- Return the structured result

## Entity

- RateLimiter
- RateLimiterAlgorithm [TokenBucketAlgorithm, SlidingWindowAlgorithm]
- ConfigurationStore
- Configuration
- Request
- Result

Relationships:

```code
    RateLimiter <----- composed of --- RateLimiterAlgorithm
    RateLimiter <----- contains ---- ConfigurationStore
    
```

## Class Desing

```code
Class RateLimiter:
    - configurationStore: ConfigurationStore
    - rateLimiterAlgorithm: AlgorithmFactory
    - states: Map<Key, RateLimiterAlgorithm>

    + registerEndpoints(endpoint string, endpintConfig EndpointConfig)
    + check(clientID: string, request: Request) -> Result
```

```code
Class Key:

    - clientID: string
    - endpoint: string
```

```code
Class ConfigurationStore:
    - defualt: EndpointConfig
    - endpointConfigs: Map<String, EndpointConfig>

    addEndpoint(endpoint: string, endpointConfig: EndpointConfig)
    getEndpointConfig(endpoint string) -> EndpointConfig
```

```code
Class AlgorithmFactory

    + create(config: EndpointConfig) -> RateLimiterAlgorithm
```

```code
Interface Class RateLimiterAlgorithm

    + allow(now: time) -> Result
```

```code
Class TokenBucketAlgorithm implements RateLimiterAlgorithm
    - configs

    + allow(now: time) -> Result
```

```code
Class SlidingWindowAlgorithm implements RateLimiterAlgorithm
    - configs
    
    + allow(now: time) -> Result
```
