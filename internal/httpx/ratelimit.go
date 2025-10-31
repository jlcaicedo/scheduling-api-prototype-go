package httpx

import (
	"net"
	"net/http"
	"sync"
	"time"
)

// simple per-IP token bucket limiter (demo only)
type bucket struct {
	avail float64
	last  time.Time
	mu    sync.Mutex
}

type limiter struct {
	rate  float64 // tokens per second
	burst float64 // max tokens
	mu    sync.Mutex
	m     map[string]*bucket
}

func newLimiter(rate float64, burst int) *limiter {
	return &limiter{
		rate:  rate,
		burst: float64(burst),
		m:     make(map[string]*bucket),
	}
}

func (l *limiter) allow(key string) bool {
	l.mu.Lock()
	b, ok := l.m[key]
	if !ok {
		b = &bucket{avail: l.burst, last: time.Now()}
		l.m[key] = b
	}
	l.mu.Unlock()

	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.last).Seconds()
	b.last = now
	// refill
	b.avail = min(l.burst, b.avail+elapsed*l.rate)
	if b.avail >= 1 {
		b.avail -= 1
		return true
	}
	return false
}

func RateLimit(rate float64, burst int) Middleware {
	l := newLimiter(rate, burst)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			if ip == "" {
				ip = "unknown"
			}
			if !l.allow(ip) {
				Error(r.Context(), w, http.StatusTooManyRequests, "rate_limited", "too many requests")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
