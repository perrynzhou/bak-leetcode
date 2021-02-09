package main

// link : https://leetcode-cn.com/problems/valid-parentheses/
import (
	"fmt"
)

func isValid(s string) bool {
  if len(s)<=1 {
  	return false
  }
  stack :=[]string{}
  mp :=map[string]string {
  "(":")",
  "{":"}",
  "[":"]",
  }
  var match int
  for i:=0;i<len(s);i++ {
  	if s[i]=='{'||s[i]=='['||s[i]=='(' {
  		stack = append(stack,string(s[i]))
  		continue
	}
	stackLen := len(stack)
	if stackLen > 0 {
		key := stack[stackLen-1]
		if mp[key]!=string(s[i]) {
			return false
		}
		stack =stack[:stackLen-1]

		match++
	}
  }
  if len(s)!=(2*match) || (len(stack)>0 && match<=0) {
  	return false
  }
  return true
}

func main() {
	fmt.Println(isValid("(({{}}})"))
}
