package main

import (
	"fmt"
	"math"
)

// link:https://leetcode-cn.com/problems/first-missing-positive/
func firstMissingPositive(nums []int) int {

	fmt.Println("nums:", nums)
	dict := make(map[int]struct{})
	for _, val := range nums {
		if val > 0 {
			dict[val] = struct{}{}
		}
	}
	fmt.Println("dict:", dict)
	for i := 1; i <= len(nums)+1; i++ {
		if _, ok := dict[i]; !ok {
			fmt.Println("i=", i)
			return i
		}
	}
	return 1
}
func firstMissingPositive_v1(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		index := int(math.Abs(float64(nums[i])))
		if index <= n {
			j := index - 1
			nums[j] = -int(math.Abs(float64(nums[j])))
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
func main() {
	/*
		fmt.Println(firstMissingPositive([]int{1,2,0}))

		fmt.Println(firstMissingPositive([]int{3,4,-1,1}))
		fmt.Println(firstMissingPositive([]int{7,8,9,11,12}))
		fmt.Println(firstMissingPositive([]int{1}))


	*/
	fmt.Println(firstMissingPositive_v1([]int{1, 2, 0}))

	fmt.Println(firstMissingPositive_v1([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive_v1([]int{7, 8, 9, 11, 12}))

	//fmt.Println(firstMissingPositive_v1([]int{7,8,9,11,12}))
	//fmt.Println(firstMissingPositive_v1([]int{1}))

}
