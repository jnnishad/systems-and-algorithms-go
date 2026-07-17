package concurrency

import (
	"errors"
	"testing"
)

func TestWorkerPool_RunsAllJobsAndPreservesIndex(t *testing.T) {
	pool := NewWorkerPool(4)

	jobs := make([]Job, 10)
	for i := 0; i < 10; i++ {
		i := i
		jobs[i] = func() (any, error) { return i * i, nil }
	}

	results := pool.Run(jobs)

	if len(results) != 10 {
		t.Fatalf("expected 10 results, got %d", len(results))
	}
	for i, r := range results {
		if r.Index != i {
			t.Errorf("result %d has wrong index %d", i, r.Index)
		}
		if r.Err != nil {
			t.Errorf("result %d unexpected error: %v", i, r.Err)
		}
		if r.Value != i*i {
			t.Errorf("result %d: expected %d, got %v", i, i*i, r.Value)
		}
	}
}

func TestWorkerPool_CapturesJobErrors(t *testing.T) {
	pool := NewWorkerPool(2)

	wantErr := errors.New("boom")
	jobs := []Job{
		func() (any, error) { return nil, wantErr },
		func() (any, error) { return "ok", nil },
	}

	results := pool.Run(jobs)

	if results[0].Err != wantErr {
		t.Errorf("expected job 0 to return the sentinel error, got %v", results[0].Err)
	}
	if results[1].Value != "ok" {
		t.Errorf("expected job 1 to succeed with 'ok', got %v", results[1].Value)
	}
}

func TestWorkerPool_RecoversPanics(t *testing.T) {
	pool := NewWorkerPool(1)

	jobs := []Job{
		func() (any, error) { panic("kaboom") },
	}

	results := pool.Run(jobs)

	if results[0].Err == nil {
		t.Fatal("expected panic to be converted into an error")
	}
}

func TestWorkerPool_EmptyJobList(t *testing.T) {
	pool := NewWorkerPool(3)
	results := pool.Run(nil)
	if len(results) != 0 {
		t.Fatalf("expected no results for empty job list, got %d", len(results))
	}
}
