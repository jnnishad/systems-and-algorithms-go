package concurrency

import (
	"sync"
	"time"
)

// TokenBucket is a classic token-bucket rate limiter: tokens refill
// continuously at `refillRate` per second up to `capacity`, and each
// Allow() call consumes one token if available. This is the same
// primitive behind most API gateway / ingress rate limiting, and the
// one most infra engineers end up hand-rolling at least once.
type TokenBucket struct {
	mu         sync.Mutex
	capacity   float64
	tokens     float64
	refillRate float64 // tokens per second
	lastRefill time.Time
	now        func() time.Time // overridable for tests
}

func NewTokenBucket(capacity float64, refillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
		now:        time.Now,
	}
}

// Allow consumes one token and reports whether the request may proceed.
func (b *TokenBucket) Allow() bool {
	return b.AllowN(1)
}

// AllowN consumes n tokens atomically — either all n are available and
// are consumed, or none are and the call is rejected.
func (b *TokenBucket) AllowN(n float64) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.refill()

	if b.tokens >= n {
		b.tokens -= n
		return true
	}
	return false
}

func (b *TokenBucket) refill() {
	now := b.now()
	elapsed := now.Sub(b.lastRefill).Seconds()
	if elapsed <= 0 {
		return
	}
	b.tokens += elapsed * b.refillRate
	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}
	b.lastRefill = now
}
