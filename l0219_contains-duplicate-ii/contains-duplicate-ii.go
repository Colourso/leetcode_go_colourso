package l0219_contains_duplicate_ii

// 给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/contains-duplicate-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func containsNearbyDuplicate(nums []int, k int) bool {
	// 暴力 + 哈希
	numMap := make(map[int]int, 0)
	for index, num := range nums {
		if value, ok := numMap[num]; ok {
			if index-value <= k {
				return true
			} else {
				numMap[num] = index
			}
		}
		numMap[num] = index
	}
	return false
}
