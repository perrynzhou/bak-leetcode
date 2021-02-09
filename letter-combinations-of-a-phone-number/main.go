package main

import "strings"

func letterCombinations(digits string) []string {
	mp := map[int][]string{
		2:{"d", "e", "f"},
		3:{"a", "b", "c"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6:{"m", "n", "o"},
		7:{"p", "q", "r", "s"},
		8:{"t", "u", "v"},
		9:{"w", "x", "y", "z"},
	}

	result :=make([]string,0)
	if len(digits) ==0 ||len(digits)>4  {
		return result
	}
	for i:=0;i<len(digits);i++ {
		if digits[i]<'2'||digits[i]>'9' {
			return result
		}
	}
	
}
