// Package dp contains dynamic-programming problems, each with the
// recurrence explained in a comment — the part that actually matters
// in an interview, more than the final code.
package dp

// CoinChange returns the fewest number of coins from `coins` needed to
// make `amount`, or -1 if it's not possible.
//
// Bottom-up DP: dp[i] = fewest coins to make amount i. Recurrence:
// dp[i] = 1 + min(dp[i-c]) over every coin c <= i. O(amount * len(coins))
// time, O(amount) space — the standard "unbounded knapsack" shape.
func CoinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}

	const unreachable = 1 << 30
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = unreachable
	}

	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if c <= 0 || c > i {
				continue
			}
			if dp[i-c]+1 < dp[i] {
				dp[i] = dp[i-c] + 1
			}
		}
	}

	if dp[amount] == unreachable {
		return -1
	}
	return dp[amount]
}
