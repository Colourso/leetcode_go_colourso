package l0040_combination_sum_ii

import "sort"

/*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。

https://leetcode.cn/problems/combination-sum-ii/description/
*/

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	sum, path, result := 0, []int{}, [][]int{}
	backtracking(candidates, target, sum, 0, path, &result)
	return result
}

func backtracking(candidates []int, target int, sum int, startIndex int, path []int, result *[][]int) {
	if sum > target {
		return
	}

	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		// 元素重复，但是整体结果不能重复，在回溯树的同一层级做过滤
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}

		path = append(path, candidates[i])
		sum += candidates[i]
		backtracking(candidates, target, sum, i+1, path, result)
		sum -= candidates[i]
		path = path[:len(path)-1]
	}
}
