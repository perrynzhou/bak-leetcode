package main

import "fmt"

//link:https://leetcode-cn.com/problems/missing-number/
func missingNumber(nums []int) int {
	var result int
	dict := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		dict[nums[i]] = struct{}{}
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := dict[i]; !ok {
			result = i
			break
		}
	}
	return result
}
func missingNumber_v2(nums []int) int {
	var result int
	for i := 0; i <= len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] == i {
				count++
				fmt.Printf("nums[%d]=%d\n", j, nums[j])
				break
			}
		}
		if count == 0 {
			result = i
			fmt.Println("result:", result)
			break
		}
	}
	return result
}
func main() {
	missingNumber_v2([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
}
