package main

import (
	"fmt"
	"strconv"
)

//link:https://leetcode-cn.com/problems/add-binary/
func addBinary_v2(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	for i := len(a) - len(b) - 1; i >= 0; i-- {
		b = "0" + b
	}
	var result string
	var flag int
	fmt.Println("a:", a, ",b:", b)
	for i := len(a) - 1; i >= 0; i-- {
		d1 := int(a[i] - '0')
		d2 := int(b[i] - '0')
		sum := d1 + d2 + flag
		if sum == 2 {
			flag = 1
			result = "0" + result
		} else if sum == 3 {
			flag = 1
			result = "1" + result
		} else {
			result = strconv.Itoa(sum) + result
			flag = 0
		}
	}
	if flag == 1 {
		result = "1" + result
	}
	return result
}
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
func main() {
	fmt.Println(addBinary("1", "11"))
	fmt.Println(addBinary("110", "11"))

}
