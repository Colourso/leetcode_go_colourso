package l0383_ransom_note

/*
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。

https://leetcode.cn/problems/ransom-note/description/
*/

func canConstruct(ransomNote string, magazine string) bool {
	// 由于只有小写字母，除了用map以外还可以用固定长度的数组来做
	chars := [26]int{}

	for i := range magazine {
		chars[magazine[i]-'a']++
	}

	for j := range ransomNote {
		if chars[ransomNote[j]-'a'] == 0 {
			return false
		}
		chars[ransomNote[j]-'a']--
	}
	return true
}
