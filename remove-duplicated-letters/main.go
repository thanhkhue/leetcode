package main

/*
https://leetcode.com/problems/remove-duplicate-letters/

Input: s = "bcabc"
Output: "abc"

Input: s = "cbacdcbc"
Output: "acdb"

*/

func main() {
	// s := "bcabc"
	// removeDuplicateLetters(s)

	s1 := "bcabc"
	removeDuplicateLetters(s1)
}

func removeByIndex(s string, i int) string {
	return s[:i] + s[i:]
}

func removeDuplicateLetters(s string) string {
	if len(s) == 0 {
		return s
	}

	appearances := make([]int, 26)
	for i := 0; i < len(s); i++ {
		appearances[s[i]-'a']++
	}

	visited := make([]int, 26)
	res := ""
	for i := 0; i < len(s); i++ {
		appearances[s[i]-'a']--
		if visited[s[i]-'a'] == 0 {
			for len(res) > 0 && s[i] < res[len(res)-1] && appearances[res[len(res)-1]-'a'] > 0 {
				visited[res[len(res)-1]-'a'] = 0
				res = res[:len(res)-1]
			}
			res += string(s[i])
			visited[s[i]-'a'] = 1
		}
	}

	return res
}
