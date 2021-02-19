package main

import "fmt"

//link :https://leetcode-cn.com/problems/first-unique-character-in-a-string/
func firstUniqChar(s string) int {
	index := -1
	dict := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if dict[s[i]] == 1 {
			index = i
			break
		}
	}
	return index
}
func firstUniqChar_v2(s string) int {
	index := -1
	dict := [26]int{}
	for i := 0; i < len(dict); i++ {
		dict[i] = -1
	}
	for i := 0; i < len(s); i++ {
		dict[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if dict[s[i]-'a'] == 0 {
			index = i
			break
		}
	}
	return index
}
func main() {
	/*
		fmt.Println(firstUniqChar("loveleetcode"))
		fmt.Println(firstUniqChar("leetcode"))
		fmt.Println(firstUniqChar("cc"))


	*/
	fmt.Println(firstUniqChar_v2("loveleetcode"))
	fmt.Println(firstUniqChar_v2("leetcode"))
	fmt.Println(firstUniqChar_v2("cc"))

}
