package l0209_minimum_size_subarray_sum

// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/minimum-size-subarray-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func minSubArrayLen(target int, nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	// 滑动窗口
	minLen, sum := -1, 0
	// i:滑动窗口起始位置, j:滑动窗口结束位置
	i := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			subLen := j - i + 1
			if minLen == -1 || subLen < minLen {
				minLen = subLen
			}
			sum -= nums[i]
			i++
		}
	}
	if minLen == -1 {
		return 0
	}
	return minLen
}

// 滑动窗口题型
// 循环的索引表示滑动窗口终止位置，
// 滑动窗口起始位置需要另一个index表示，index走过的位置已经没有了计算价值
// 小序列【b,c】符合结果，那么大序列【a,b,c】也符合结果，目标是求最小的，
//所以变动的是起始窗口位置，最终窗口位置作为for循环条件，保证能够遍历整个循环。
