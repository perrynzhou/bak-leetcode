package main

import "fmt"

//link:https://leetcode-cn.com/problems/counting-bits/

func countBits_v1(num int) []int {
	var result []int
	result = append(result, 0)
	for i := 1; i <= num; i++ {
		tmp := i
		count := 0
		for {
			if tmp == 0 {
				break
			}
			if tmp%2 == 1 {
				count++
			}
			tmp = tmp / 2
		}
		result = append(result, count)
	}
	fmt.Println("result:", result)
	return result
}
func countBits(num int) []int {
	result := make([]int, num+1)
	result[0] = 0
	for i := 1; i <= num; i++ {
		fmt.Println("value:", i, "i&i-1=", i&(i-1))
		result[i] = result[i&(i-1)] + 1
	}
	fmt.Println("result:", result)
	return result
}

func main() {
	countBits_v1(5)
	countBits(5)

}
