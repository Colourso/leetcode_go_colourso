package l0343_integer_break

/*
给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。

返回 你可以获得的最大乘积 。

https://leetcode.cn/problems/integer-break/description/
*/

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = getMax(dp[i], getMax(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
