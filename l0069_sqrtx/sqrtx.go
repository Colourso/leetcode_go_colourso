package l0069_sqrtx

// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
//
//由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
//
//注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/sqrtx
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 暴力：
// ans目标值的查找范围是(0,x)
// 就是一个一个整数进行尝试
// for i:=0;i<x;i++ {
//	if i*i <= x && (i+1)*(i+1)>x {
//		return i //即找到了目标值
//	}
//}

func mySqrt(x int) int {
	// 边界条件负数与0
	if x <= 0 {
		return 0
	}
	left, right := 0, x
	// 边界条件1
	for left <= right {
		mid := (left + right) / 2
		if mid*mid <= x && (mid+1)*(mid+1) > x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
