package l0015_3sum

import "sort"

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，

同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

https://leetcode.cn/problems/3sum/description/
*/

func threeSum(nums []int) [][]int {
	results := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			// 最小值大于0时，加和一定大于0
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			// 重复搜索过的可去重
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				results = append(results, []int{nums[i], nums[l], nums[r]})
				l++
				r--

				for l < r && nums[l] == nums[l-1] {
					l++
				}

				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return results
}
