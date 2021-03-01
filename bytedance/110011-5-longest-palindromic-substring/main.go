package main

import "fmt"

//link: https://leetcode-cn.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	fmt.Println("origin:", s)
	left, right := []byte{}, []byte{}
	n := len(s)
	if n < 1 {
		return string(left)
	}
	m, n := 0, 0
	for ; m < len(s); m++ {
		n = m + 1
		for n < len(s) {
			if s[m] == s[n] {
				break
			}
			n++
		}
	}
	fmt.Println("s:", s[m:n])
	i, j := m, n
	for i < j {
		if s[i] == s[j] {
			left = append(left, s[i])
			right = append(right, s[j])
		}
		i++
		j--
	}
	fmt.Println("left:", string(left), "right:", string(right))
	if i == j {
		left = append(left, s[i])
	}
	for i := 0; i < len(right); i++ {
		left = append(left, right[i])
	}
	return string(left)
}
func main() {
	fmt.Println(longestPalindrome("babad"))
	//cbbd
	fmt.Println(longestPalindrome("cbbd"))

}
