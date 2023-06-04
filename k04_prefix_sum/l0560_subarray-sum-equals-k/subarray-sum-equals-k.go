package l0560_subarray_sum_equals_k

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。
//
//
//
//示例 1：
//
//输入：nums = [1,1,1], k = 2
//输出：2
//示例 2：
//
//输入：nums = [1,2,3], k = 3
//输出：2
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/subarray-sum-equals-k
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func subarraySum(nums []int, k int) int {
	// 滑动窗口: 全为正数时可以使用
	// 但 0,0,0 1,-1,0这种情况下，滑动窗口会因满足条件，而收缩，从而减少可能，这时需要使用前缀和思想

	// 前缀和之差 就是连续子数组的和
	preSumMap := make(map[int]int)
	preSumMap[0] = 1 // 前缀和0，出现了1次

	preSum, count := 0, 0

	// 边计算，边合并
	for _, num := range nums {
		preSum += num
		// prei - prej = k; prej = prei - k
		if value, ok := preSumMap[preSum-k]; ok {
			count += value
		}
		preSumMap[preSum] += 1
	}

	return count
}

// 时间复杂度 o(n)
