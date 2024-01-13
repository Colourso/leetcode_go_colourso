package l0027_remove_elemen

import (
	"fmt"
	"testing"
)

func TestRemoveElementV2(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	lens := removeElementV2(nums, val)
	fmt.Println(lens)

	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	lens = removeElementV2(nums2, 2)
	fmt.Println(lens)
}
