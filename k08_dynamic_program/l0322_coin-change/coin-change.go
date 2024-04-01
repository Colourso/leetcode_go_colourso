package l0322_coin_change

import (
	"fmt"
	"math"
)

/**
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。

https://leetcode.cn/problems/coin-change/description/
*/

func coinChange(coins []int, amount int) int {
	// dp【i】凑成最少的种类数
	// 最少种类数
	// dp[i] -- dp[i] 小 还是 dp[i-j]+1小
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := range coins {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = getMin(dp[j], dp[j-coins[i]]+1)
				fmt.Println(i, j, dp[j])
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func getMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
