package l0077_combinations

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。

https://leetcode.cn/problems/combinations/description/
*/

func combine(n int, k int) [][]int {
	var path []int
	var result [][]int
	backtracking(n, k, 1, path, &result)
	return result
}

func backtracking(n int, k int, startIndex int, path []int, result *[][]int) {
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	for i := startIndex; i <= n; i++ {
		path = append(path, i)
		backtracking(n, k, i+1, path, result)
		path = path[:len(path)-1]
	}
}
