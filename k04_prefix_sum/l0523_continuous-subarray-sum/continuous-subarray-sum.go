package l0523_continuous_subarray_sum

// 给你一个整数数组 nums 和一个整数 k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：
//
//子数组大小 至少为 2 ，且
//子数组元素总和为 k 的倍数。
//如果存在，返回 true ；否则，返回 false 。
//
//如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。0 始终视为 k 的一个倍数。
//
//
//
//示例 1：
//
//输入：nums = [23,2,4,6,7], k = 6
//输出：true
//解释：[2,4] 是一个大小为 2 的子数组，并且和为 6 。
//示例 2：
//
//输入：nums = [23,2,6,4,7], k = 6
//输出：true
//解释：[23, 2, 6, 4, 7] 是大小为 5 的子数组，并且和为 42 。
//42 是 6 的倍数，因为 42 = 7 * 6 且 7 是一个整数。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/continuous-subarray-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func checkSubarraySum(nums []int, k int) bool {
	// 连续子数组
	preSumMap := make(map[int]int)
	preSumMap[0] = -1

	// 如何保证子数组大小为2，
	// 如何保证子数组之和为k的倍数，如何匹配
	// 1. 同余定理，(m-n)%k == 0 --> m%k == n%k
	// 2. map存储，加快查找，key:余数，value:下标，使用下标来计算子数组大小超过2
	// 3. map覆盖？一个是同余定理，只要匹配就行，另一个是for循环加深，先出现的余数就能保证子数组大小会很大，这就是最大的情况

	preSum, remind := 0, 0
	for index, num := range nums {
		preSum += num
		remind = preSum % k
		if preIndex, ok := preSumMap[remind]; ok {
			if index-preIndex >= 2 {
				return true
			}
		} else {
			preSumMap[remind] = index
		}
	}
	return false
}
