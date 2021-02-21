package main

import (
	"fmt"
	"strconv"
)

//link:https://leetcode-cn.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x <0 {
		return false
	}
  str := strconv.Itoa(x)
  result := true
  l,r := 0,len(str)-1
  for l<=r {
  	  if str[l]!=str[r] {
  	  	result = false
  	  	break
	  }
	  l++
	  r--
  }
  return result
}
func main() {
	fmt.Println()
}
