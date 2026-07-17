package dp

// LongestCommonSubsequence returns the length of the longest
// subsequence common to both a and b (not necessarily contiguous).
//
// Classic 2D DP: dp[i][j] = LCS length of a[:i] and b[:j].
// If a[i-1] == b[j-1], that character extends the LCS found for the
// prefixes without it: dp[i][j] = dp[i-1][j-1] + 1. Otherwise it's the
// best of dropping one character from either string. O(len(a)*len(b))
// time and space.
func LongestCommonSubsequence(a, b string) int {
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
