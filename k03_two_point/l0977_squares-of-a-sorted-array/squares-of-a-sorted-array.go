package l0977_squares_of_a_sorted_array

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，
// 要求也按 非递减顺序 排序。
// https://leetcode.cn/problems/squares-of-a-sorted-array/

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	// 非递减数组，外加有负数，那么左右两边的平方结果都是要大于中间的
	// 故使用双指针

	left, right, index := 0, len(nums)-1, len(nums)-1
	for left < right {
		if nums[left]*nums[left] >= nums[right]*nums[right] {
			result[index] = nums[left] * nums[left]
			left++
		} else {
			result[index] = nums[right] * nums[right]
			right--
		}
		index--
	}
	result[index] = nums[left] * nums[left]
	return result
}
