package l0062_unique_paths_ii

import "testing"

func TestName(t *testing.T) {
	dp := [][]int{
		{0, 1},
		{0, 0}}

	uniquePathsWithObstacles(dp)
}
