// Package arrays contains classic array/string interview problems,
// each solved at the optimal known complexity with the reasoning for
// why in a comment above the function.
package arrays

// TwoSum returns the indices of the two numbers in nums that add up to
// target, or (0, 0, false) if no such pair exists.
//
// O(n) time / O(n) space: a single pass with a value->index map, instead
// of the O(n^2) brute-force nested loop.
func TwoSum(nums []int, target int) (int, int, bool) {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		complement := target - n
		if j, ok := seen[complement]; ok {
			return j, i, true
		}
		seen[n] = i
	}
	return 0, 0, false
}
