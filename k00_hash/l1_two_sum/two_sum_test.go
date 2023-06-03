package l1_two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	// nums = [3,3], target = 6 输出：[0,1]
	ans1, ans2 := []int{0, 1}, []int{1, 0}
	if ans := TwoSum([]int{3, 3}, 6); !reflect.DeepEqual(ans, ans1) && !reflect.DeepEqual(ans, ans2) {
		t.Errorf("nums = [3,3], target = 6 输出：[0,1], but error")
	}

	// go语言 二维数组 的地址连续吗？

}
