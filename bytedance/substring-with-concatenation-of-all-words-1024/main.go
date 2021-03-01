package main

import "fmt"

//link:https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/
func findSubstring(s string, words []string) []int {
	result := []int{}
	dict := map[string]int{}
	arrs := make([]string, len(words))

	for i := range words {
		dict[words[i]] = 0
		arrs[i] = words[i]
	}

	for i := 0; i < len(s); i++ {
		for m := 0; m < len(words); m++ {
			k := 0
			p := i
			for key, val := range dict {
				//fmt.Println("1===")
				if p >= len(s) {
					break
				}
				if val == 0 && s[p] == key[k] {
					break
				}
				//	fmt.Println("2===")

				k = len(key)
				//fmt.Println("p:",p,"k:",k,",p+k=",p+k,",s len:",len(s))
				if p+k > len(s) {
					break
				}
				fmt.Println("key:", key, ",s[p+k]:", s[p:p+k])

				if s[p:p+k] == key {
					dict[key] = 1
					fmt.Println("key:", key, ",s:", s[p:p+k])
				}
				k = 0
				p += len(key)
			}
			for key, _ := range dict {
				dict[key] = 0
			}
		}

	}
	return result
}
func main() {
	//findSubstring("barfoothefoobarman",[]string{"foo","bar"})
	//findSubstring("wordgoodgoodgoodbestword",[]string{"word","good","best","word"})
	//"wordgoodgoodgoodbestword"
	//["word","good","best","good"]
	//findSubstring("wordgoodgoodgoodbestword",[]string{"word","good","best","good"})
	//"lingmindraboofooowingdingbarrwingmonkeypoundcake"
	//["fooo","barr","wing","ding","wing"]
	findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"})

}
