package L0088_merge_sorted_array

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

https://leetcode.cn/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 直接赋值，然后排序 --o(n^2)
	// 本身有序了，主要做的是移动元素 -- 存储到新空间之中
	result := make([]int, m+n)
	p1, p2, cur := 0, 0, 0
	for p1 < m || p2 < n {
		if p1 < m && p2 < n {
			if nums1[p1] <= nums2[p2] {
				result[cur] = nums1[p1]
				p1++
			} else {
				result[cur] = nums2[p2]
				p2++
			}
		} else if p1 < m {
			result[cur] = nums1[p1]
			p1++
		} else {
			result[cur] = nums2[p2]
			p2++
		}
		cur++
	}
	nums1 = result
	return
}
