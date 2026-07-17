package concurrency

import (
	"testing"
	"time"
)

func TestTokenBucket_AllowsUpToCapacity(t *testing.T) {
	b := NewTokenBucket(3, 1) // 3 tokens, refill 1/sec

	for i := 0; i < 3; i++ {
		if !b.Allow() {
			t.Fatalf("expected token %d to be allowed", i)
		}
	}
	if b.Allow() {
		t.Fatal("expected 4th immediate request to be rejected — bucket should be empty")
	}
}

// fixedClock pins both the bucket's "now" and its internal lastRefill
// to the same instant, so the elapsed-time math in refill() is exact
// instead of off by however many nanoseconds elapsed between test
// statements.
func fixedClock(b *TokenBucket, at time.Time) {
	b.now = func() time.Time { return at }
	b.lastRefill = at
}

func TestTokenBucket_RefillsOverTime(t *testing.T) {
	start := time.Now()
	b := NewTokenBucket(1, 1) // 1 token, refills at 1/sec
	fixedClock(b, start)

	if !b.Allow() {
		t.Fatal("expected first request to be allowed")
	}
	if b.Allow() {
		t.Fatal("expected immediate second request to be rejected")
	}

	// Advance the fake clock by exactly 1 second — exactly one token refills.
	b.now = func() time.Time { return start.Add(1 * time.Second) }
	if !b.Allow() {
		t.Fatal("expected request to be allowed after refill window elapsed")
	}
}

func TestTokenBucket_AllowNRejectsPartialConsumption(t *testing.T) {
	start := time.Now()
	b := NewTokenBucket(5, 0)
	fixedClock(b, start)

	if !b.AllowN(3) {
		t.Fatal("expected 3 tokens to be available")
	}
	// 2 tokens remain; asking for 3 more should fail entirely, not
	// partially consume.
	if b.AllowN(3) {
		t.Fatal("expected AllowN(3) to fail when only 2 tokens remain")
	}
	if !b.AllowN(2) {
		t.Fatal("expected the remaining 2 tokens to still be available")
	}
}
