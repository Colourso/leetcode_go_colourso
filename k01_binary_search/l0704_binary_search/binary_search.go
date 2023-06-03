package l0704_binary_search

// 704. 二分查找
// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
// 写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/binary-search
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	var left, right = 0, len(nums) - 1
	for left <= right {
		mid := left + ((right - left) / 2)
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 二分查找：无重复数据，且数据升序排序时
// 每一次使用mid值和target作比较，然后能缩减一半的数据比较
// 时间复杂度计算：n表示数据规模~
//	数组长度为n，那么最坏找到一个数据需要查多少次呢？即循环代码执行多少次呢？设为x次，2^x = n,
// 则时间复杂度是o(logn)
// 空间复杂度，由于只使用了常数个变量，故空间复杂度为o(1)
