package l1_two_sum

// TwoSum
// https://leetcode.cn/problems/two-sum/
// 使用哈希表来做处理
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, n := range nums {
		if key1, ok := numMap[target-n]; ok {
			return []int{key1, index}
		}
		numMap[n] = index
	}
	return nil
}
