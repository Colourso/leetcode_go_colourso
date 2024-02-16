package l0078_subsets

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

https://leetcode.cn/problems/subsets/description/
*/

func subsets(nums []int) [][]int {
	result := [][]int{}
	var backtracking func(startIndex int, path []int)
	backtracking = func(startIndex int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)

		if startIndex >= len(nums) {
			return
		}

		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(i+1, path)
			path = path[:len(path)-1]
		}
	}
	backtracking(0, []int{})
	return result
}
