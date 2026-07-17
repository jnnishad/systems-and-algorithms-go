package dp

import "testing"

func TestCoinChange(t *testing.T) {
	cases := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{"classic example", []int{1, 2, 5}, 11, 3}, // 5+5+1
		{"exact single coin", []int{2}, 4, 2},
		{"impossible", []int{2}, 3, -1},
		{"zero amount needs zero coins", []int{1, 2, 5}, 0, 0},
		{"no coins available", []int{}, 7, -1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CoinChange(tc.coins, tc.amount)
			if got != tc.want {
				t.Fatalf("CoinChange(%v, %d) = %d, want %d", tc.coins, tc.amount, got, tc.want)
			}
		})
	}
}
