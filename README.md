# Scheduling API Prototype — Go / REST (2025)

Proyecto de ejemplo con una API REST ligera usando **Go** y solo la librería estándar (`net/http`). Incluye:

- Middleware de autenticación (Bearer) y **rate limiting**.
- Respuestas JSON estructuradas y trazas por **Request-ID**.
- Endpoints de ejemplo para *health* y *schedules*.
- Integración de prueba con un cliente **Swift** (URLSession).
- Dockerfile y CI básico (Go 1.22+).

> Objetivo: servir como base pública para GitHub que demuestre prácticas sólidas y un código claro.

## Requisitos

- Go 1.22+
- (Opcional) Docker
- (Opcional) Make

## Ejecutar

```bash
# Variables de entorno (token de ejemplo)
export API_BEARER_TOKEN="dev-secret-token"   # en producción, cámbialo
export API_ADDR=":8080"

# Ejecuta
go run ./cmd/api
```

Prueba rápida:

```bash
# Health (no requiere auth)
curl -s http://localhost:8080/health | jq

# Listar schedules (requiere auth)
curl -s http://localhost:8080/v1/schedules   -H "Authorization: Bearer $API_BEARER_TOKEN" | jq

# Crear schedule (requiere auth)
curl -s -X POST http://localhost:8080/v1/schedules   -H "Authorization: Bearer $API_BEARER_TOKEN"   -H "Content-Type: application/json"   -d '{"title":"Demo","time":"2025-11-01T10:00:00Z"}' | jq
```

## Estructura

```
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── httpx/
│   │   ├── middleware.go
│   │   ├── ratelimit.go
│   │   └── respond.go
│   └── schedules/
│       ├── handler.go
│       └── store.go
├── go.mod
├── Makefile
├── Dockerfile
├── .gitignore
└── .github/workflows/ci.yml
```

## Diseño

### Respuestas JSON
Se usa un sobre (envelope) consistente:
```json
{
  "status": "ok",
  "data": {...},
  "error": null,
  "meta": {"request_id":"..."}
}
```

### Autenticación
- Esquema **Bearer** vía cabecera `Authorization`.  
- Token esperado en `API_BEARER_TOKEN` (solo para demo). En producción, integrar un *provider* real (OAuth2/JWT).

### Rate limiting
- Límite por IP con *token bucket* en memoria (solo para demo).  
- Config por defecto: 5 reqs/seg, *burst* 10. Ajustable por env:
  - `RL_RATE_PER_SEC` (float, p.ej. `5`)
  - `RL_BURST` (int, p.ej. `10`)

### Integración Swift (cliente demo)
```swift
import Foundation

struct Schedule: Codable { let id: String; let title: String; let time: String }

let baseURL = URL(string: "http://localhost:8080")!
let token = ProcessInfo.processInfo.environment["API_BEARER_TOKEN"] ?? "dev-secret-token"

var req = URLRequest(url: baseURL.appendingPathComponent("/v1/schedules"))
req.httpMethod = "GET"
req.addValue("Bearer \(token)", forHTTPHeaderField: "Authorization")

URLSession.shared.dataTask(with: req) { data, resp, err in
    if let data = data {
        let decoder = JSONDecoder()
        if let obj = try? decoder.decode([String: [Schedule]].self, from: data) {
            print(obj)
        } else {
            print(String(data: data, encoding: .utf8) ?? "no data")
        }
    } else {
        print(err ?? "unknown error")
    }
}.resume()
```

## Pruebas mínimas
```bash
go test ./...
```

## Licencia
MIT
