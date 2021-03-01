package main

import (
	"fmt"
)

//link:https://leetcode-cn.com/problems/add-binary/

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	res := make([]byte, len(a))
	j, tmp := len(b)-1, byte(0)
	for i := len(a) - 1; i >= 0; i-- {
		if j >= 0 {
			tmp += b[j] - '0'
			j--
		}
		tmp += a[i] - '0'
		res[i] = (tmp % 2) + '0'
		tmp /= 2
	}
	if tmp > 0 {
		return "1" + string(res)
	}
	return string(res)
}

func addBinary_v2(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	result := make([]byte, len(a))
	j, tmp := len(b)-1, byte(0)
	for i := len(a) - 1; i >= 0; i-- {
		if j >= 0 {
			tmp += b[j] - '0'
			j--
		}
		tmp += a[i] - '0'
		result[i] = (tmp % 2) + '0'
		tmp /= 2
	}
	if tmp > 0 {
		return "1" + string(result)
	}
	return string(result)
}
func main() {
	fmt.Println(addBinary_v2("1", "11"))
	fmt.Println(addBinary_v2("110", "11"))

}
