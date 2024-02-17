package l0090_subsets_ii

import "sort"

/*
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

https://leetcode.cn/problems/subsets-ii/description/
*/

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	var backtracking func(startIndex int, path []int)
	backtracking = func(startIndex int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)

		for i := startIndex; i < len(nums); i++ {
			if i > startIndex && nums[i] == nums[i-1] {
				// 横向遍历表示树的扩展，这种情况下和前一个元素想同则证明被使用过
				continue
			}
			path = append(path, nums[i])
			backtracking(i+1, path)
			path = path[:len(path)-1]
		}

	}

	backtracking(0, []int{})
	return result
}
