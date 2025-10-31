# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o /api ./cmd/api

FROM scratch
ENV API_ADDR=":8080"
COPY --from=build /api /api
EXPOSE 8080
ENTRYPOINT ["/api"]
