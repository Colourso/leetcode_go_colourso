package l0974_subarray_sums_divisible_by_k

// 给定一个整数数组 nums 和一个整数 k ，返回其中元素之和可被 k 整除的（连续、非空） 子数组 的数目。
//
//子数组 是数组的 连续 部分。
//
//
//
//示例 1：
//
//输入：nums = [4,5,0,-2,-3,1], k = 5
//输出：7
//解释：
//有 7 个子数组满足其元素之和可被 k = 5 整除：
//[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/subarray-sums-divisible-by-k
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 1.连续子数组之和，即前缀和之差
// 2.元素之和可被k整除，

func subarraysDivByK(nums []int, k int) int {
	if len(nums) <= 0 {
		return 0
	}

	prefixSumMap := map[int]int{0: 1}
	prefixSum, count := 0, 0

	for _, num := range nums {
		prefixSum += num
		remind := (prefixSum%k + k) % k
		count += prefixSumMap[remind]
		prefixSumMap[remind]++
	}

	return count
}
