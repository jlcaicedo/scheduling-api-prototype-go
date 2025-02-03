# Worklog — Scheduling API Prototype — Go / REST (2025)

Registro de trabajo del prototipo REST en Go (net/http), con bearer auth, rate limit, Request-ID y ejemplos de schedules + cliente Swift.


## Feb–Mar 2025 (base REST, middleware, JSON & Swift demo)

- Arranque con net/http, middleware (auth + rate limit), envelope JSON, Request-ID, endpoints health/schedules y cliente Swift.

## Abr–Jun 2025 (Docker/CI, hardening, observabilidad y refinamientos)

- Dockerfile/CI, endurecimiento de endpoints, performance/observabilidad y pulido DX.
- 2025-02-01: Go: project scaffold (cmd/api, internal/...)
- 2025-02-01: net/http: basic mux & routing
- 2025-02-02: Middleware: bearer auth header parser
- 2025-02-02: Middleware: rate limiting (token bucket in-memory)
- 2025-02-02: Request-ID: context propagation & logger
- 2025-02-03: JSON: envelope {status,data,error,meta}
