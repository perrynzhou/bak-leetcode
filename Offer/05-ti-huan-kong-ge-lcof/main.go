package main

import "fmt"

func replaceSpace(s string) string {
	var result []byte
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			result = append(result, s[i])
		} else {
			result = append(result, '%', '2', '0')
		}
	}
	return string(result)
}
func main() {
	fmt.Println(replaceSpace("We are happy."))
}
