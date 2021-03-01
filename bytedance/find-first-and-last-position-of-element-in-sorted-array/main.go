package main

import "fmt"

// link:https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	index := []int{-1, -1}
	if nums == nil {
		return index
	}
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == target {
			if index[0] == -1 {
				index[0] = i
				index[1] = i
			} else {
				if i > index[1] {
					index[1] = i
				}
			}
		}
		i++
	}
	fmt.Println("index:", index)
	return index
}
func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 5, 5, 5}, 5))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 10))

}
