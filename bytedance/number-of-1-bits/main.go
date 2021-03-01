package main

import (
	"fmt"
	"strconv"
)

// link:https://leetcode-cn.com/problems/number-of-1-bits/
func hammingWeight(num uint32) int {

	s := strconv.FormatUint(uint64(num), 2)
	fmt.Println("value:", s)

	l, r := 0, len(s)-1
	if l == r && s[l] == '1' {
		return 1
	}
	count := 0
	for l < r {
		//fmt.Println("s:",string(s[0]))
		if s[l] == '1' {
			count++
		}
		if s[r] == '1' {
			count++
		}
		l++
		r--
	}
	if l == r && s[l] == '1' {
		count++
	}
	fmt.Println("count:", count)
	return count
}

func main() {
	hammingWeight(01010101010101011010101010101010)

}
