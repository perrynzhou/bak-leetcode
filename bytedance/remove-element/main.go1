package main

import (
	"fmt"
)

// link:https://leetcode-cn.com/problems/remove-element/
// [3,2,2,3]
func removeElement(nums []int, val int) int {
	if nums == nil  {
		return 0
	}
	var j int
	var index int
	for i:=0;i<len(nums);i++ {
		if nums[i]==val {
			j++
		}
		if nums[i]!=val {
			nums[index]=nums[i]
			index++
		}
	}
	if j>0 {
		nums=nums[:len(nums)-j]
	}
	fmt.Println("Data:",nums)
   return len(nums)
}
func main() {
  fmt.Println(removeElement([]int{3,2,1,1,3,3,4},3))
}
