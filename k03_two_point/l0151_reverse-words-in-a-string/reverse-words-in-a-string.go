package l0151_reverse_words_in_a_string

import "strings"

/*
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

示例 1：

输入：s = "the sky is blue"
输出："blue is sky the"

https://leetcode.cn/problems/reverse-words-in-a-string/description/
*/

// reverseWords 使用新空间处理
func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	result := ""
	charBool := false
	left, right := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if charBool {
				// 存储结果
				right = i
				result = " " + string(s[left:right]) + result
				charBool = false
			}
		}

		if s[i] != ' ' {
			if !charBool {
				charBool = true
				left = i
			}
		}
	}

	if charBool {
		result = s[left:right] + result
	} else {
		if len(result) > 0 {
			result = result[1:]
		}
	}

	return result
}

// reverseWordsV2 不使用额外空间
func reverseWordsV2(s string) string {
	str := trimSpace(s)
	reverse(str, 0, len(str)-1)
	reverseAllWord(str)
	return string(str)
}

// trimSpace 去除首尾空格，单词中空格仅保留一个
func trimSpace(s string) []byte {
	left, right := 0, len(s)-1
	for left < right && s[left] == ' ' {
		left++
	}
	for left < right && s[right] == ' ' {
		right--
	}

	str := make([]byte, 0, right-left+1)
	// 注意 left = right 有意义
	for left <= right {
		if s[left] != ' ' {
			str = append(str, s[left])
		} else if str[len(str)-1] != ' ' {
			str = append(str, s[left])
		}
		left++
	}
	return str
}

// 翻转数组
func reverse(str []byte, left int, right int) {
	for left < right {
		str[left], str[right] = str[right], str[left]

		left++
		right--
	}
}

// reverseAllWord 再翻转每个数组里的单词
func reverseAllWord(str []byte) {
	left, right, n := 0, 0, len(str)
	for left < n {
		for right < n && str[right] != ' ' {
			right++
		}
		reverse(str, left, right-1)
		left, right = right+1, right+1
	}
}
