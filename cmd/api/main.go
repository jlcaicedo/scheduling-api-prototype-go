package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/your-user/scheduling-api-prototype-go/internal/httpx"
	"github.com/your-user/scheduling-api-prototype-go/internal/schedules"
)

func main() {
	addr := getenv("API_ADDR", ":8080")
	token := getenv("API_BEARER_TOKEN", "dev-secret-token")

	ratePerSec := atof(getenv("RL_RATE_PER_SEC", "5"))
	burst := atoi(getenv("RL_BURST", "10"))

	mux := http.NewServeMux()

	// health: no auth
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		httpx.OK(r.Context(), w, map[string]string{"status": "up"})
	})

	// in-memory store
	store := schedules.NewStore()

	// protected routes under /v1
	mux.Handle("GET /v1/schedules", httpx.Chain(
		schedules.ListHandler(store),
		httpx.AuthBearer(token),
		httpx.RateLimit(ratePerSec, burst),
		httpx.RequestID(),
		httpx.Recoverer(),
	))

	mux.Handle("POST /v1/schedules", httpx.Chain(
		schedules.CreateHandler(store),
		httpx.AuthBearer(token),
		httpx.RateLimit(ratePerSec, burst),
		httpx.RequestID(),
		httpx.Recoverer(),
	))

	srv := &http.Server{
		Addr:              addr,
		Handler:           httpx.WithCommon(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	// graceful shutdown
	idleConnsClosed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server shutdown error: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	log.Printf("listening on %s\n", printableAddr(addr))
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
	<-idleConnsClosed
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func atof(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func printableAddr(addr string) string {
	if addr == "" {
		return ":http"
	}
	_, port, err := net.SplitHostPort(addr)
	if err != nil || port == "" {
		return addr
	}
	return "http://localhost:" + port
}
