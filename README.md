# systems-and-algorithms-go

Algorithms and data structures in Go, split into two halves on purpose:

- **Classic DSA** (`arrays/`, `linkedlist/`, `trees/`, `graphs/`, `dp/`) — the
  interview-loop staples, each solved at optimal complexity with the
  reasoning written above the function, not just the code.
- **Systems patterns** (`concurrency/`) — the building blocks that
  actually show up in infra tooling: an LRU cache, a token-bucket rate
  limiter, a worker pool. These come up far more often in an SRE/infra
  interview than reversing a linked list does, and they're the part
  most DSA-practice repos skip.

## Why this exists

Most of [this profile](https://github.com/jnnishad) is infrastructure —
Terraform, Kubernetes, Ansible. None of it demonstrates comfort with
core data structures and algorithmic complexity, which is a real part
of most senior infra/SRE interview loops (Google's included). This repo
is that half: same standard as everything else — real tests, complexity
noted, edge cases covered — applied to DSA instead of infra config.

## Structure

```
arrays/         two-sum, sliding-window-max
linkedlist/     reverse, cycle detection (Floyd's)
trees/          BST validation, level-order (BFS) traversal
graphs/         topological sort — the actual algorithm behind
                sequencing Terraform module applies or Helm chart
                dependencies, tested against a GitOps-shaped example
dp/             coin change, longest common subsequence
concurrency/    LRU cache, token-bucket rate limiter, worker pool —
                all goroutine-safe, all tested with -race
```

## Notes on approach

Every solution has:

1. A comment explaining **why** this approach and complexity, not just
   what the code does.
2. Table-driven tests covering the standard case, edge cases (empty
   input, single element), and — where the bug risk is real, like
   `sliding_window_max` — a property-based fuzz test against a naive
   brute-force reference implementation.
3. No unnecessary dependencies: this module imports nothing outside
   the Go standard library.

## Run it

```bash
go test ./... -race -cover
go vet ./...
```

## Status

Written without a local Go toolchain (hand-traced instead of compiled),
then verified for real via GitHub Actions on push. First CI run caught
one bug: a wrong test fixture in `validate_bst_test.go` (my example
tree wasn't actually invalid) — fixed, all packages now pass
`go test ./... -race -cover`.

<!-- test commit 2026-02-02T13:28:57 -->

<!-- test commit 2026-03-28T16:37:09 -->

<!-- test commit 2026-04-20T17:11:50 -->

<!-- test commit 2026-05-08T11:11:22 -->
