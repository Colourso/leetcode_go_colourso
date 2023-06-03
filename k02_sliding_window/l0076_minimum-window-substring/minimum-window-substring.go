package l0076_minimum_window_substring

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//
//
//注意：
//
//对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
//如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//示例 1：
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/minimum-window-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func minWindow(s string, t string) string {
	result := ""
	if len(s) < len(t) {
		return result
	}

	// go没有set集合
	// go字符串中的每个字符都是byte类型？
	tCharSets := make(map[byte]int) // 等价于? tCharSets := map[byte]int{}
	for i := 0; i < len(t); i++ {
		tCharSets[t[i]]++
	}

	sCharTmpSet := make(map[byte]int)

	i := 0
	// go滑动窗口
	for j := 0; j < len(s); j++ {
		sCharTmpSet[s[j]]++

		// 符合条件时，缩小
		for sCharContainTChar(sCharTmpSet, tCharSets) {
			if result == "" || len(result) > j-i+1 {
				result = s[i : j+1] // 切片规则
			}
			sCharTmpSet[s[i]]--
			i++
		}
	}
	return result
}

func sCharContainTChar(sCharTmpSet, tCharSets map[byte]int) bool {
	for key := range tCharSets {
		if sCharTmpSet[key] < tCharSets[key] {
			return false
		}
	}
	return true
}
