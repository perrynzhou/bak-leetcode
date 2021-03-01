package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	var result int
	tmp := strings.TrimSpace(s)
	if len(tmp) == 0 {
		return result
	}
	sign := 1
	var index int
	if tmp[index] == '-' || tmp[index] == '+' {
		if tmp[index] == '-' {
			sign = -1
		}
		index++
	}
	if index == len(tmp) {
		return result
	}
	if tmp[index] < '0' || tmp[index] > '9' {
		return result
	}
	for i := index; i < len(tmp); i++ {
		if tmp[i] < '0' || tmp[i] > '9' {
			break
		}
		value := int(tmp[i] - '0')
		result = result*10 + value
		if sign*result > math.MaxInt32 {
			result = math.MaxInt32
			return result
		}
		if sign*result < math.MinInt32 {
			result = math.MinInt32
			return result
		}
	}
	return sign * result
}
func main() {
	fmt.Println(myAtoi("-91283472332"))
}
