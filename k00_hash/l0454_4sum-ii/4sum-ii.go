package l0454_4sum_ii

/*
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：
0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

https://leetcode.cn/problems/4sum-ii/description/
*/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sumMap, count := make(map[int]int), 0

	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			sumMap[n1+n2]++
		}
	}

	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			count += sumMap[-n3-n4]
		}
	}
	return count
}
