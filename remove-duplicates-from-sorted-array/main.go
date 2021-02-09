package main

import "fmt"

// link:https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) <= 1 {
		return len(nums)
	}
	var index int
	nums[index] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[index] != nums[i] {
			index++
			nums[index] = nums[i]
		}
		if i == len(nums)-1 {
			index++
		}
		fmt.Println("nums:", nums, ",index:", index)
	}
	if index > 0 {
		nums = nums[:index]
	}
	fmt.Println(nums)
	return len(nums)
}
func main() {
	//0,0,1,1,1,2,2,3,3,4
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4,5,5,5,5,7,7,9,9}))
	fmt.Println(removeDuplicates([]int{1, 2}))
	fmt.Println(removeDuplicates([]int{1, 2, 2}))
	fmt.Println(removeDuplicates([]int{1, 1}))

}
