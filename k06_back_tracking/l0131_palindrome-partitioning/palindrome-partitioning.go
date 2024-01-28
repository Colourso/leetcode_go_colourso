package l0131_palindrome_partitioning

/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。

https://leetcode.cn/problems/palindrome-partitioning/description/
*/

func partition(s string) [][]string {
	path, result := []string{}, [][]string{}
	backtracking(s, 0, path, &result)
	return result
}

func backtracking(s string, startIndex int, path []string, result *[][]string) {
	if startIndex >= len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	for i := startIndex; i < len(s); i++ {
		if isPalindrome(s, startIndex, i) {
			path = append(path, s[startIndex:i+1])
		} else {
			continue
		}
		backtracking(s, i+1, path, result)
		path = path[:len(path)-1]
	}
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
