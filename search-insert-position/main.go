package main

//link: https://leetcode-cn.com/problems/search-insert-position/submissions/
func searchInsert(nums []int, target int) int {
	index := -1
	i, j := 0, len(nums)-1
	if target < nums[i] {
		index = i
		return index
	}
	if target > nums[j] {
		index = j + 1
		return index
	}
	prevIndex := -1
	for i <= j {
		if nums[i] == target {
			index = i
			break
		}
		if nums[i] > target {
			prevIndex = i
			break
		}
		i++
	}
	if index == -1 {
		index = prevIndex
	}
	return index
}
