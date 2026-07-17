package arrays

// SlidingWindowMax returns the maximum of every contiguous window of
// size k as nums slides across it.
//
// O(n) time using a monotonic decreasing deque of indices — each index
// enters and leaves the deque at most once — instead of the naive
// O(n*k) approach of scanning every window from scratch.
func SlidingWindowMax(nums []int, k int) []int {
	if k <= 0 || len(nums) == 0 {
		return nil
	}
	if k > len(nums) {
		k = len(nums)
	}

	var deque []int // stores indices into nums, values strictly decreasing
	result := make([]int, 0, len(nums)-k+1)

	for i, n := range nums {
		// Drop indices that fall out of the current window.
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		// Maintain the decreasing invariant: anything smaller than the
		// new element can never be the max again while n is in range.
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= n {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)

		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
