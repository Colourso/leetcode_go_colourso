package k046_beibao01

import (
	"fmt"
)

/*
小明是一位科学家，他需要参加一场重要的国际科学大会，以展示自己的最新研究成果。他需要带一些研究材料，但是他的行李箱空间有限。
这些研究材料包括实验设备、文献资料和实验样本等等，它们各自占据不同的空间，并且具有不同的价值。

小明的行李空间为 N，问小明应该如何抉择，才能携带最大价值的研究材料，每种研究材料只能选择一次，并且只有选与不选两种选择，不能进行切割。

输入描述
第一行包含两个正整数，第一个整数 M 代表研究材料的种类，第二个正整数 N，代表小明的行李空间。

第二行包含 M 个正整数，代表每种研究材料的所占空间。

第三行包含 M 个正整数，代表每种研究材料的价值。

输出描述
输出一个整数，代表小明能够携带的研究材料的最大价值。

提示信息
小明能够携带 6 种研究材料，但是行李空间只有 1，而占用空间为 1 的研究材料价值为 5，所以最终答案输出 5。

数据范围：
1 <= N <= 5000
1 <= M <= 5000
研究材料占用空间和价值都小于等于 1000

*/

func Beibao01() int {
	M, N := 0, 0
	fmt.Scanln(&M, &N)
	weight := make([]int, M)
	value := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&weight[i])
	}
	for i := 0; i < M; i++ {
		fmt.Scan(&value[i])
	}

	dp := make([][]int, M)
	for i := range dp {
		dp[i] = make([]int, N+1) // 容量0也要计算
	}
	for i := 0; i < N; i++ {
		if weight[0] <= i {
			dp[0][i] = value[0]
		}
	}

	// 先物品、后背包
	for i := 1; i < M; i++ {
		for j := 1; j <= N; j++ {
			if weight[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = getMax(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}

	return dp[M-1][N]
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Beibao01V2() int {
	M, N := 0, 0
	fmt.Scanln(&M, &N)
	weight := make([]int, M)
	value := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&weight[i])
	}
	for i := 0; i < M; i++ {
		fmt.Scan(&value[i])
	}

	dp := make([]int, N+1)
	for i := range dp {
		dp[i] = 0
	}

	// 先物品、后背包：先看物品0在背包产生的最大价值，然后看物品1累计后的，一直继续...
	// 滚动数组，使用一维数组
	// 为什么要倒序遍历呢？上一层的 dp[j] 到了这一层 dp[j]
	// 递推数组要求是【依赖前一个数据】（那么前一个数据就要最后更新），正序遍历的话前一个数据提前被更新了，就达不到依赖的效果了
	for i := 0; i < M; i++ {
		for j := N; j >= weight[i]; j-- {
			dp[j] = getMax(dp[j], dp[j-weight[i]]+value[i])
		}
	}

	return dp[N]
}
