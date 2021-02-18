package main

import (
	"fmt"
	"math"
)

// link:https://leetcode-cn.com/problems/divide-two-integers/
func divide(dividend int, divisor int) int {
	fmt.Println("step2-dividend:", dividend)
	if divisor == 1 || divisor == -1 {
		if dividend*divisor > math.MaxInt32 {
			return math.MaxInt32
		}
		if dividend*divisor < math.MinInt32 {
			return math.MinInt32
		}
		return dividend * divisor
	}
	sign := 1
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		sign = -1
	}
	fmt.Println("step2-sign:", sign)

	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))
	var result int
	for {
		tmp := dividend - divisor
		if tmp < 0 {
			break
		}
		result++
		dividend -= divisor
	}
	if sign*result > math.MaxInt32 {
		return math.MaxInt32
	}
	if sign*result < math.MinInt32 {
		return math.MinInt32
	}
	return sign * result
}
func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(-2147483648, -1))

}
