package l0074_search_a_2d_matrix

// 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/search-a-2d-matrix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func SearchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left < right {
		mid := (left + right) / 2
		row, col := mid/n, mid%n
		if matrix[row][col] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// left == right 时停止
	return target == matrix[left/n][left%n]
}
