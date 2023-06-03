package l0074_search_a_2d_matrix

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	// go 二维数组如何初始化
	target1 := 9
	target2 := 10

	find1 := SearchMatrix(matrix, target1)

	find2 := SearchMatrix(matrix, target2)

	fmt.Printf("%v - %v", find1, find2)

}
