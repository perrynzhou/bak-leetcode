package main

import "fmt"

// link:https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates_ans(nums []int) int {
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

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) <= 1 {
		return len(nums)
	}
	var index int
	//[1,1,1,2,2,3,4]

	for i := 0; i < len(nums); i++ {
		j := i + 1
		if j == len(nums) {
			break
		}
		if nums[i] == nums[j] {
			if index == 0 {
				index = j
			}
		} else {
			if index > 0 {
				nums[index] = nums[j]
				index++
			}
		}
	}
	if index > 0 {
		nums = nums[:index]

	}
	fmt.Println(nums)

	return len(nums)
}
func removeDuplicates_ans2(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) <= 1 {
		return len(nums)
	}
	fmt.Println("origin:", nums)
	var index int
	//[1,1,1,2,2,3,4]
	nums[index] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[index] != nums[i] {
			index++
			nums[index] = nums[i]
		}
		if index > 0 && i == len(nums)-1 {
			index++
		}
		fmt.Println("i:", i, "nums:", nums, ",index:", index)
	}
	if index > 0 {
		nums = nums[:index]
	}
	if index == 0 {
		nums = nums[:1]
	}
	fmt.Println(nums)
	return len(nums)
}

func main() {
	//0,0,1,1,1,2,2,3,3,4
	fmt.Println(removeDuplicates([]int{0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 4, 5, 5, 5, 5}))
	fmt.Println(removeDuplicates([]int{1, 2, 2}))
	fmt.Println(removeDuplicates([]int{1, 1}))

	/*
		fmt.Println(removeDuplicates([]int{1, 2}))
		fmt.Println(removeDuplicates([]int{1, 2, 2}))
		fmt.Println(removeDuplicates([]int{1, 1}))

	*/

}
