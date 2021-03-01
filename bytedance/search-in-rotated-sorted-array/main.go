package main

//link:https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	index := -1
	l, r := 0, len(nums)-1
	if target < nums[l] && target > nums[r] {
		return index
	}
	for l <= (len(nums)-1) && r >= 0 {
		if nums[l] == target {
			index = l
			break
		}
		if nums[r] == target {
			index = r
			break
		}
		l++
		r--
	}
	return index
}
