package arrays

import "testing"

func TestTwoSum(t *testing.T) {
	cases := []struct {
		name       string
		nums       []int
		target     int
		wantFound  bool
		wantValues [2]int // expected nums[i], nums[j], order-independent
	}{
		{"simple match", []int{2, 7, 11, 15}, 9, true, [2]int{2, 7}},
		{"match at end", []int{3, 2, 4}, 6, true, [2]int{2, 4}},
		{"duplicate values", []int{3, 3}, 6, true, [2]int{3, 3}},
		{"no match", []int{1, 2, 3}, 100, false, [2]int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			i, j, found := TwoSum(tc.nums, tc.target)
			if found != tc.wantFound {
				t.Fatalf("expected found=%v, got %v", tc.wantFound, found)
			}
			if !found {
				return
			}
			if i == j {
				t.Fatalf("expected two distinct indices, got the same index twice: %d", i)
			}
			gotSum := tc.nums[i] + tc.nums[j]
			if gotSum != tc.target {
				t.Fatalf("nums[%d]+nums[%d] = %d, want %d", i, j, gotSum, tc.target)
			}
		})
	}
}
