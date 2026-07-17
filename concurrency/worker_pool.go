package concurrency

import "sync"

// WorkerPool runs jobs across a fixed number of goroutines and collects
// their results, preserving no particular order but guaranteeing every
// job's result (or panic-recovered error) is returned exactly once.
type WorkerPool struct {
	workers int
}

func NewWorkerPool(workers int) *WorkerPool {
	if workers <= 0 {
		workers = 1
	}
	return &WorkerPool{workers: workers}
}

// Job is a unit of work that returns a result or an error.
type Job func() (any, error)

// Result pairs a job's index (so callers can correlate results back to
// input) with its outcome.
type Result struct {
	Index int
	Value any
	Err   error
}

// Run executes all jobs across the pool's fixed worker count and
// returns once every job has completed. A panic inside a job is
// recovered and surfaced as an error rather than crashing the pool.
func (p *WorkerPool) Run(jobs []Job) []Result {
	results := make([]Result, len(jobs))

	type indexed struct {
		index int
		job   Job
	}

	work := make(chan indexed)
	var wg sync.WaitGroup

	for w := 0; w < p.workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range work {
				results[item.index] = runSafely(item.index, item.job)
			}
		}()
	}

	for i, j := range jobs {
		work <- indexed{index: i, job: j}
	}
	close(work)

	wg.Wait()
	return results
}

func runSafely(index int, job Job) (result Result) {
	defer func() {
		if r := recover(); r != nil {
			result = Result{Index: index, Err: panicError{r}}
		}
	}()

	value, err := job()
	return Result{Index: index, Value: value, Err: err}
}

type panicError struct{ v any }

func (p panicError) Error() string {
	return "job panicked: " + toString(p.v)
}

func toString(v any) string {
	if err, ok := v.(error); ok {
		return err.Error()
	}
	if s, ok := v.(string); ok {
		return s
	}
	return "non-string panic value"
}
