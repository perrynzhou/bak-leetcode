package main

import (
	"fmt"
)

//link : https://leetcode-cn.com/problems/longest-common-prefix/
func cmpString(i int, strs []string) bool {
	ok := true
	for k, _ := range strs {
		if k < len(strs)-1 {
			if strs[k][:i+1] != strs[k+1][:i+1] {
				ok = false
				break
			}
		}
	}
	return ok
}
func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	var result string
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		n := len(strs[i])
		if minLen >= n {
			minLen = n
		}
	}
	for i := 0; i < minLen; i++ {
		if cmpString(i, strs) {
			result = strs[0][:i+1]
		}
	}

	return result
}
func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//["dog","racecar","car"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
