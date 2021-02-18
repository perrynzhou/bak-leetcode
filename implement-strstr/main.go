package main

import "fmt"

//link:https://leetcode-cn.com/problems/implement-strstr/

func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	fmt.Println("l1:", l1, ",l2:", l2)

	if l1 == 0 && l2 != 0 {
		return -1
	}
	if l1 == 0 && l2 == l1 {
		return 0
	}
	if l2 > l1 {
		return -1
	}
	if l1 == l2 && haystack == needle {
		return 0
	}
	for i := 0; i < l1; i++ {
		j := 0
		if j < l2 && haystack[i] != needle[j] {
			continue
		}
		j = l2
		if i+j > l1 {
			break
		}
		//fmt.Println("##j:",j,"i:",i,",haystack:",haystack[i:i+j],",needle:",needle)

		if haystack[i:i+j] == needle {
			//fmt.Println("##i:",i,",haystack:",haystack[i:i+j],",needle:",needle)
			return i
		}
	}
	return -1
}
func main() {
	//fmt.Println("index:",strStr("hello","ll"))
	//fmt.Println("index:",strStr("aaaaaa","bba"))
	//fmt.Println("index:",strStr("abc","c"))
	//fmt.Println("index:",strStr("mississippi","issip"))

	fmt.Println("index:", strStr("", "*"))

	//fmt.Println("index:",strStr("*",""))

}
