package main

import (
	"net/http"
	"sync"
	"time"
)

// token bucket limiter

type rateLimiter struct {
	rate  int
	burst int
	token int
	time  time.Time
	mu    sync.Mutex
}

type chanrateLimiter struct {
	tokens chan struct{}
}

func NewRateLimiter(rate, burst int) *rateLimiter {
	return &rateLimiter{
		rate:  rate,
		burst: burst,
		token: burst,
		time:  time.Now(),
	}
}

func (rl *rateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(rl.time).Seconds()
	rl.token += int(elapsed * float64(rl.rate))
	if rl.rate > rl.burst {
		rl.token = rl.burst
	}
	rl.time = now
	if rl.token > 0 {
		rl.token--
		return true
	}
	return false
}

func RateLimitMiddleware(rl *rateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !rl.Allow() {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("429 - rate limit exceeded"))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func chanRateLimiter(rate int, burst int) *chanrateLimiter {
	rl := &chanrateLimiter{
		tokens: make(chan struct{}),
	}
	//fill the bucket
	for i := 0; i < burst; i++ {
		rl.tokens <- struct{}{}
	}
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rate))
		defer ticker.Stop()
		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}:
				//add the token
			default:
				// full bucket
			}
		}
	}()
	return rl

}
func (rl *chanrateLimiter) ChanAllow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func main() {
	rl := NewRateLimiter(10, 20)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	http.ListenAndServe(":8080", RateLimitMiddleware(rl)(mux))
}
