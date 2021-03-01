package main

import (
	"fmt"
	"strings"
)

// link:https://leetcode-cn.com/problems/count-and-say/

// 1
// 11
// 21
// 1211
// 111221
// 312211
func countAndSay(n int) string {
	str := "1"
	for i := 2; i <= n; i++ {
		byte := str[0]
		count := 1
		var tmp strings.Builder
		for j := 1; j < len(str); j++ {
			if byte == str[j] {
				count++
			} else {
				out := fmt.Sprintf("%d%s", count, string(byte))
				fmt.Println("out:", out)
				tmp.WriteString(out)
				count = 1
				byte = str[j]
			}
		}
		out := fmt.Sprintf("%d%s", count, string(byte))
		tmp.WriteString(out)
		str = tmp.String()

	}
	return str
}
func main() {
	fmt.Println(countAndSay(5))
}
