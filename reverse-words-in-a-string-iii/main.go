package main

import (
	"fmt"
)

//link :https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/

func reverseWords(s string) string {
	n := len(s)
	result := []byte{}
	start, end := -1, 0
	for i := 0; i < n; i++ {
		if s[i] != ' ' && start == -1 {
			start = i
			//	fmt.Printf("ch=%c,start:%d\n",s[i],start)
		}
		if s[i] == ' ' {
			end = i - 1
			//	fmt.Printf("[%d,%d]===%c,%c\n",start,end,s[start],s[end])
			for k := end; k >= start; k-- {
				result = append(result, s[k])
			}
			result = append(result, s[i])
			start = -1
		}
	}
	if start != -1 {
		for k := len(s) - 1; k >= start; k-- {
			result = append(result, s[k])
		}
	}
	//fmt.Printf("[%d,%d]===%c,%c\n",start,len(s)-1,s[start],s[len(s)-1])
	//fmt.Printf("result:%s\n",string(result))

	return string(result)
}
func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
