package main

import (
	"fmt"
	"sort"
)

//link:https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
func findRepeatNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	target := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			target = nums[i]
			break
		}
	}
	fmt.Println("target:", target)
	return target
}
func main() {
	findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3})
}
