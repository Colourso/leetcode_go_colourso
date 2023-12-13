package l1423_maximum_points_you_can_obtain_from_cards

//https://leetcode.cn/problems/maximum-points-you-can-obtain-from-cards/description/
//几张卡牌 排成一行，每张卡牌都有一个对应的点数。点数由整数数组 cardPoints 给出。
//
//每次行动，你可以从行的开头或者末尾拿一张卡牌，最终你必须正好拿 k 张卡牌。
//
//你的点数就是你拿到手中的所有卡牌的点数之和。
//
//给你一个整数数组 cardPoints 和整数 k，请你返回可以获得的最大点数。
//
//
//
//示例 1：
//
//输入：cardPoints = [1,2,3,4,5,6,1], k = 3
//输出：12
//解释：第一次行动，不管拿哪张牌，你的点数总是 1 。
//但是，先拿最右边的卡牌将会最大化你的可获得点数。
//最优策略是拿右边的三张牌，最终点数为 1 + 6 + 5 = 12 。

func maxScore(cardPoints []int, k int) int {
	sum, left, result := 0, 0, 0
	for i := range cardPoints {
		sum += cardPoints[i]
	}
	minNum := sum
	minSubLens := len(cardPoints) - k
	// 考虑长度超出的情况
	if minSubLens <= 0 {
		return sum
	}

	for right := 0; right < len(cardPoints); right++ {
		result += cardPoints[right]

		for right-left+1 == minSubLens {
			if minNum > result {
				minNum = result
			}

			result -= cardPoints[left]
			left++
		}
	}
	return sum - minNum
}

// 题解：正向看来很难模拟，每次两两比较两端选择一个最大值这是局部最优的贪心，他不是全局最优
// 全局最优需要穷举所有可能得结果，然后比较得出，模拟的代码不太好写
// 逆向来看，最大值=总和-最小值，并且最小值还是剩余列表中连续的 -- (连续的个数为len-k！)
// 解题思路就变成求连续子数组的最小值了，一个滑动窗口题目
