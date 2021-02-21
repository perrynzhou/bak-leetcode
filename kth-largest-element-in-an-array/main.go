package main

import "sort"

//link :https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
func findKthLargest(nums []int, k int) int {
   sort.Slice(nums,func(i,j int)bool {
   	return nums[i]<nums[j]
   })
   return nums[len(nums)-k]
}