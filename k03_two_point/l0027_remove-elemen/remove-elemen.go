package l0027_remove_elemen

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
//
//
//说明:
//
//为什么返回数值是整数，但输出的答案是数组呢?
//
//请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//
//你可以想象内部操作如下:
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/remove-element
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func removeElement(nums []int, val int) int {
	if len(nums) <= 0 {
		return 0
	}

	left, right := 0, len(nums)-1
	for left < right {
		// 找到了目标
		for left < right && nums[left] != val {
			left++
		}

		// 直接与右指针替换
		if left < right {
			tmp := nums[right]
			nums[right] = nums[left]
			nums[left] = tmp
			right--
		}
	}
	// 当左右指针相等的时候，右指针右边的一定是val值
	// go语言没有三言表达式
	if nums[left] == val {
		return right
	}
	return right + 1
}

// 原地移除 —> 如果一个元素匹配再移除的话，就要连续挪动后面的元素，最坏的时间复杂度接近o（n^2）
//
// 但是双指针进行替换的话，就没有这些担心,空间位置上的直接替换，没有什么损伤
