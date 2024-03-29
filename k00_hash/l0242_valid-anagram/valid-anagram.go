package l0242_valid_anagram

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

https://leetcode.cn/problems/valid-anagram/description/
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := make(map[byte]int)
	count := 0
	for i := 0; i < len(s); i++ {
		charMap[s[i]] = charMap[s[i]] + 1
		count++
	}

	for i := 0; i < len(s); i++ {
		if charMap[t[i]] > 0 {
			charMap[t[i]] = charMap[t[i]] - 1
			count--
		}
	}
	return count == 0
}
