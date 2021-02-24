package main

import (
	"strings"
)

//link: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		in := strings.Index(string(s[start:i]), string(s[i]))
		if in == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += in + 1
			end += in + 1
		}
	}
	return end - start
}
