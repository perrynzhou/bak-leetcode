package main

import (
	"fmt"
	"strings"
)

// link:https://leetcode-cn.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	ss := strings.Split(s, " ")
	if len(ss) == 0 {
		return 0
	}

	for i := len(ss) - 1; i >= 0; i-- {
		if len(ss[i]) > 0 {
			return len(ss[i])
		}
	}
	fmt.Println("ss:", ss)
	return 0
}
func main() {
	fmt.Println(lengthOfLastWord("a "))
}
