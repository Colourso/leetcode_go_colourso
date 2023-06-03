package l0034_find_first_and_last_position_of_element_in_sorted_array

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。
//请你找出给定目标值在数组中的开始位置和结束位置。
//
//如果数组中不存在目标值 target，返回 [-1, -1]。
//
//你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func SearchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}

	// 1,搜索第一个出现的位置
	result[0] = getLeftBorder(nums, target)

	// 2,搜索最后一个出现的位置
	result[1] = getRightBorder(nums, target)

	return result
}

// getRightBorder 找到数组中的右边界 找最后一个小于等于target的数
func getRightBorder(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right + 1) / 2 // 特殊条件！！
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}

// getLeftBorder 找到数组中的左边界 找第一个大于等于target的数
func getLeftBorder(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// 当搜索停止的时候 left == right
	if nums[left] == target {
		return left
	}

	return -1
}
