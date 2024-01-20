package l0018_4sum

import "sort"

/*
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。

https://leetcode.cn/problems/4sum/description/
*/

func fourSum(nums []int, target int) [][]int {
	results := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, len(nums)-1
			for l < r {
				if nums[i]+nums[j]+nums[l]+nums[r] == target {
					results = append(results, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--

					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if nums[i]+nums[j]+nums[l]+nums[r] > target {
					r--
				} else {
					l++
				}
			}
		}
	}

	return results
}
