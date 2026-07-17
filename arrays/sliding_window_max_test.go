package arrays

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSlidingWindowMax(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"classic example", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{"window size 1 returns input", []int{4, 2, 9}, 1, []int{4, 2, 9}},
		{"window size equals length", []int{1, 2, 3}, 3, []int{3}},
		{"empty input", []int{}, 3, nil},
		{"descending sequence", []int{5, 4, 3, 2, 1}, 2, []int{5, 4, 3, 2}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := SlidingWindowMax(tc.nums, tc.k)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("SlidingWindowMax(%v, %d) = %v, want %v", tc.nums, tc.k, got, tc.want)
			}
		})
	}
}

// TestSlidingWindowMax_MatchesBruteForce fuzzes the optimized deque
// implementation against a naive O(n*k) reference on random input,
// which catches off-by-one window-boundary bugs that hand-picked cases
// tend to miss.
func TestSlidingWindowMax_MatchesBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(42))

	for trial := 0; trial < 200; trial++ {
		n := rng.Intn(20) + 1
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(41) - 20
		}
		k := rng.Intn(n) + 1

		got := SlidingWindowMax(nums, k)
		want := bruteForceMax(nums, k)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("trial %d: nums=%v k=%d: got %v, want %v", trial, nums, k, got, want)
		}
	}
}

func bruteForceMax(nums []int, k int) []int {
	if k <= 0 || len(nums) == 0 {
		return nil
	}
	if k > len(nums) {
		k = len(nums)
	}
	result := make([]int, 0, len(nums)-k+1)
	for i := 0; i+k <= len(nums); i++ {
		max := nums[i]
		for j := i + 1; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		result = append(result, max)
	}
	return result
}
