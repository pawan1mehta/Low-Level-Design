# Rate Limiter

Link: <https://www.hellointerview.com/learn/low-level-design/problem-breakdowns/rate-limiter>

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
- Configuration is provided at the satrtup
  - Algorithm to use
  - Algorithm specific paramters
- System enforces rate limits by checking the clientId and endpoint
- If endpoint has no configuration se the default configuration
- Return the structured result

## Entity

- RateLimiter
- Algorithm [TokenBucket, SlidingWindow]
- ConfigurationStore
- Configuration
- Request
- Result

Relationships:

```code

```

## Class Desing

```code
Class RateLimiter:
    - configurationStore: ConfigurationStore
    - rateLimiterAlgorithm: RateLimiterAlgorithm

    + registerEndpoints(clientId string, endpintConfig EndpointConfig)
    + request(request: Request) -> Result
```

```code
Class ConfigurationStore:

    - clientsConfig: Map<ClientID, Map<String, EndpointConfig>>

    addEndpoint(clientID: string, endpointConfig: EndpointConfig)
    getEndpointConfig(clientId string, endpoint string) -> EndpointConfig
```

```code
Abstract Class RateLimiterAlgorithm

    + request(request: Request) -> Result
```

```code
Class TokenBucket implements RateLimiterAlgorithm

    - bucketInfo:

    + request(request: Request) -> Result
```

```code
Class SlidingWindow implements RateLimiterAlgorithm

    - windoInfo:

    + request(request: Request) -> Result
```
