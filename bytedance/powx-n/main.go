package main

import (
	"fmt"
)

//link:https://leetcode-cn.com/problems/powx-n/

func myPow_v2(x float64, n int) float64 {
	res := float64(1)
	f := 0
	if n < 0 {
		f = 1 // 确定符号
		n = -n
	}
	for n != 0 {
		if n&1 == 1 {
			res *= x // 当n为奇数时多乘以一个x
		}
		n >>= 1
		fmt.Println("n=", n)
		x *= x // 每次算平方
	}
	if f == 0 {
		return res
	}
	return 1 / res // 根据符号得到
}
func myPow(x float64, n int) float64 {
	var sign float64
	if n == 0 {
		return 1
	}
	if n < 0 {
		sign = 1
		n = -n
	}
	sum := float64(1)
	for n != 0 {
		if n&1 == 1 {
			sum *= x
		}
		n >>= 1
		x *= x
	}
	if sign == 1 {
		sum = sign / sum
	}
	return sum
}
func main() {
	/*
		fmt.Println(myPow(2.00000, 10))
		fmt.Println(myPow(2.10000, 3))
		//2.00000
		fmt.Println(myPow(2.00000, -2))

	*/

	fmt.Println(myPow_v2(2.00000, 10))
	fmt.Println(myPow_v2(2.10000, 3))
	//2.00000
	fmt.Println(myPow_v2(2.00000, -2))

	fmt.Println(myPow(0.00001, 2147483647))
	/*

		fmt.Println(myPow(0.00001, 2147483647))
		fmt.Println(myPow(2, 31))



	*/
}
