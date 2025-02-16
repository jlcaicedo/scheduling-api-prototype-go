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
- 2025-02-03: Health endpoint: /health (no auth)
- 2025-02-06: Schedules: GET /v1/schedules (auth)
- 2025-02-06: Schedules: POST /v1/schedules (auth)
- 2025-02-07: Store: in-memory schedules + IDs
- 2025-02-07: httpx: respond helpers (errors, meta)
- 2025-02-08: Config: env vars (API_ADDR, API_BEARER_TOKEN)
- 2025-02-08: Logging: structured fields (req_id, method, path)
- 2025-02-08: Tests: minimal http handlers
- 2025-02-09: Swift client: URLSession GET /v1/schedules (demo)
- 2025-02-09: Docs: README quick-start & curl snippets
- 2025-02-10: Go: project scaffold (cmd/api, internal/...)
- 2025-02-14: net/http: basic mux & routing
- 2025-02-14: Middleware: bearer auth header parser
- 2025-02-15: Middleware: rate limiting (token bucket in-memory)
- 2025-02-15: Request-ID: context propagation & logger
- 2025-02-16: JSON: envelope {status,data,error,meta}
- 2025-02-16: Health endpoint: /health (no auth)
- 2025-02-16: Schedules: GET /v1/schedules (auth)
