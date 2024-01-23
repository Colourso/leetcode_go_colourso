package l0216_combination_sum_iii

/*
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：

只使用数字1到9
每个数字 最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。

https://leetcode.cn/problems/combination-sum-iii/description/
*/

func combinationSum3(k int, n int) [][]int {
	path, result := []int{}, [][]int{}
	backtracking(k, n, 0, 1, path, &result)
	return result
}

func backtracking(k, n, sum, startIndex int, path []int, result *[][]int) {
	if len(path) == k {
		if sum == n {
			tmp := make([]int, k)
			copy(tmp, path)
			*result = append(*result, tmp)
		}
		return
	}

	for i := startIndex; i <= 9 && k-len(path)-1 <= 9-i && sum < n; i++ {
		path = append(path, i)
		sum += i
		backtracking(k, n, sum, i+1, path, result)
		sum -= i
		path = path[:len(path)-1]
	}
}
