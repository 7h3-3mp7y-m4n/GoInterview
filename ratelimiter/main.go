package main

import (
	"fmt"
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

func main() {
	rl := NewRateLimiter(5, 5)
	for i := 0; i < 20; i++ {
		allowed := rl.Allow()
		fmt.Printf("%02d: allowed=%v (tokens=%d)\n", i, allowed, rl.token)
		time.Sleep(1000 * time.Millisecond)
	}
}
