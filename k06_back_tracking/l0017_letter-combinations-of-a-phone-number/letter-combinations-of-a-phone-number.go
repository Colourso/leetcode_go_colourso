package l0017_letter_combinations_of_a_phone_number

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
*/

var (
	letters []string
)

func letterCombinations(digits string) []string {
	if len(digits) <= 0 {
		return []string{}
	}

	letters = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path, result := []byte{}, []string{}
	backtracking(digits, 0, path, &result)
	return result
}

func backtracking(digits string, startIndex int, path []byte, result *[]string) {
	if len(path) == len(digits) {
		*result = append(*result, string(path))
		return
	}

	// 为什么startIndex没有报数组越界？因为他的位置和当前path的长度是一致的，path满足后提前返回不会走进去
	chars := letters[digits[startIndex]-'0']
	for i := 0; i < len(chars); i++ {
		path = append(path, chars[i])
		backtracking(digits, startIndex+1, path, result)
		path = path[:len(path)-1]
	}
}
