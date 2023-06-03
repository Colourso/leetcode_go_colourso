package l0035_search_insert_position

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//请必须使用时间复杂度为 O(log n) 的算法。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/search-insert-position
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1
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

	return right + 1 // 此时，right < left，然后right指针指向的元素时小于target，其后一个元素又大于target
}

// 搜索插入位置，暴力法：遍历即可，O(n)
// 二分法，O(logn)
// 与二分不同的点，插入位置，搜索时如果存在，则能够找到其位置
// 如果找不到时，left、right指针怎么停止的？
