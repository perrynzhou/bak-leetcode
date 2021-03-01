package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	var result int
	s := strconv.Itoa(x)
	sign := 1
	var index int
	if s[index] == '-' {
		sign = -1
		index++
	}
	for i := len(s) - 1; i >= index; i-- {
		result = result*10 + int(s[i]-'0')
		if sign*result > math.MaxInt32 {
			return 0
		}
		if sign*result < math.MinInt32 {
			return 0
		}
	}
	return sign * result
}
func main() {
	fmt.Println(reverse(1534236469))
}
