package l0093_restore_ip_addresses

/*
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。

你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

https://leetcode.cn/problems/restore-ip-addresses/description/
*/

func restoreIpAddresses(s string) []string {
	result := []string{}
	path := []string{}
	backtracking(s, 0, path, &result)
	return result
}

func backtracking(s string, startIndex int, path []string, result *[]string) {
	if len(path) > 4 {
		return
	}

	if len(path) == 4 {
		if startIndex == len(s) {
			tmp := ""
			for i := range path {
				tmp = tmp + "." + path[i]
			}
			tmp = tmp[1:]
			*result = append(*result, tmp)
			return
		} else {
			return
		}
	}

	for i := startIndex; i < len(s); i++ {
		if isInvaild(s, startIndex, i) {
			path = append(path, s[startIndex:i+1])
			backtracking(s, i+1, path, result)
			path = path[:len(path)-1]
		} else {
			break
			// 当当前的结果不符合预期的时候，后序再怎么加都不符合预期
		}
	}

}

func isInvaild(s string, start, end int) bool {
	// 长度大于1时首位为0不符合
	if end-start+1 > 1 {
		if s[start] == '0' {
			return false
		}
	}

	// 长度大于3时不符合
	if end-start+1 > 3 {
		return false
	}

	// 整体大小大于255时不符合
	iStr := 0
	for i := range s[start : end+1] {
		iStr = iStr*10 + int(s[start+i]-'0')
	}
	if iStr > 255 {
		return false
	}

	return true
}
