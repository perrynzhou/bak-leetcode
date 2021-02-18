package main

import "fmt"

//link:https://leetcode-cn.com/problems/plus-one/

func plusOne(digits []int) []int {
	dLen := len(digits)
	result := []int{}
	var incre int
	for i := dLen - 1; i >= 0; i-- {
		if i == dLen-1 {
			digits[i] = digits[i] + 1
		}
		if digits[i] > 9 {
			incre = 1
			digits[i] = 0
		} else {
			digits[i] = digits[i] + incre
			if incre == 1 {
				incre = 0
			}
			if digits[i] > 9 {
				incre = 1
				digits[i] = 0
			}
		}
		fmt.Println("i:", i, ",nums:", digits[i], ",incre:", incre)

	}
	if incre == 1 && digits[0] == 0 {
		result = append(result, 1)
	}
	for i := 0; i < dLen; i++ {
		result = append(result, digits[i])
		fmt.Println("i:", i, ",result:", result)
	}
	return result
}
func plusOne_v2(digits []int) []int {
	dLen := len(digits)
	var incre int
	digits[dLen-1] += 1
	if digits[dLen-1] <= 9 {
		return digits
	}
	result := []int{}
	incre = 1
	digits[dLen-1] = 0
	for i := dLen - 2; i >= 0; i-- {
		digits[i] = digits[i] + incre
		if incre == 1 {
			incre = 0
		}
		if digits[i] > 9 {
			incre = 1
			digits[i] = 0
		}
		fmt.Println("i:", i, ",nums:", digits[i], ",incre:", incre)

	}
	if incre == 1 && digits[0] == 0 {
		result = append(result, 1)
	}
	for i := 0; i < dLen; i++ {
		result = append(result, digits[i])
		fmt.Println("i:", i, ",result:", result)
	}
	return result
}
func main() {

	plusOne([]int{4, 3, 2, 1})
	plusOne([]int{9, 9, 9, 9})
	plusOne([]int{1, 2, 3})

	plusOne([]int{9, 8, 9})

}
