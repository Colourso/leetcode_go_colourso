package L0088_merge_sorted_array

import "testing"

func TestName(t *testing.T) {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)

	merge([]int{0}, 0, []int{1}, 1)
}
