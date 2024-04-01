package l1013_partition_array_into_three_parts_with_equal_sum

/*
https://leetcode.cn/problems/partition-array-into-three-parts-with-equal-sum/description/

给你一个整数数组 arr，只有可以将其划分为三个和相等的 非空 部分时才返回 true，否则返回 false。

形式上，如果可以找出索引 i + 1 < j 且满足 (arr[0] + arr[1] + ... + arr[i] == arr[i + 1] + arr[i + 2] + ... + arr[j - 1] == arr[j] + arr[j + 1] + ... + arr[arr.length - 1]) 就可以将数组三等分。
*/

func canThreePartsEqualSum(arr []int) bool {
	// 连续，且总和相等的三部分
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	if sum%3 != 0 {
		return false
	}

	// 从前、从后 逼近
	i, j, k := 0, len(arr)-1, sum/3
	sum1, sum2 := 0, 0
	for ; i < j; i++ {
		sum1 += arr[i]
		if sum1 == k {
			break
		}
	}

	if i == j {
		// 此时找不到值
		return false
	}

	for ; j > i; j-- {
		sum2 += arr[j]
		if sum2 == k {
			break
		}
	}

	if j == i {
		// 此时找不到值
		return false
	}

	return true
}
