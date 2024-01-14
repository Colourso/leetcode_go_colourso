package l0349_intersection_of_two_arrays

/*
给定两个数组 `nums1` 和 `nums2` ，返回 *它们的交集* 。
输出结果中的每个元素一定是 **唯一** 的。我们可以 **不考虑输出结果的顺序** 。

https://leetcode.cn/problems/intersection-of-two-arrays/description/
*/

func intersection(nums1 []int, nums2 []int) []int {
	numsMap := make(map[int]bool)
	result := make([]int, 0)

	for i := range nums1 {
		numsMap[nums1[i]] = true
	}

	for i := range nums2 {
		unUsed, ok := numsMap[nums2[i]]
		if ok {
			if unUsed {
				result = append(result, nums2[i])
				numsMap[nums2[i]] = false
			}
		}
	}

	return result
}
