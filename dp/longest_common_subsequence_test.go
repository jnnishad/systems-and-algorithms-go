package dp

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	cases := []struct {
		name string
		a, b string
		want int
	}{
		{"classic example", "abcde", "ace", 3},
		{"no common subsequence", "abc", "def", 0},
		{"identical strings", "abc", "abc", 3},
		{"one empty string", "", "abc", 0},
		{"both empty", "", "", 0},
		{"interleaved", "AGGTAB", "GXTXAYB", 4}, // GTAB
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := LongestCommonSubsequence(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("LongestCommonSubsequence(%q, %q) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
