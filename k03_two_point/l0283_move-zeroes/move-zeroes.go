package l0283_move_zeroes

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。
https://leetcode.cn/problems/move-zeroes/description/
*/

// moveZeroes 快慢指针处理
func moveZeroes(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if 0 != nums[right] {
			nums[left] = nums[right]
			left++
		}
	}

	for left < len(nums) {
		nums[left] = 0
		left++
	}
}

// moveZeroesV2 快慢指针处理
func moveZeroesV2(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if 0 != nums[right] {
			// 不等于0的都挪到left这一侧
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
