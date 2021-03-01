package main

import "fmt"

// link:https://leetcode-cn.com/problems/house-robber/
func rob(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	if n >= 2 {
		nums[1] = max(nums[0], nums[1])
	}
	for i := 2; i < n; i++ {
		nums[i] = max(nums[i]+nums[i-2], nums[i-1])
		fmt.Println("nums:", nums)
	}
	return nums[n-1]
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {
	fmt.Println("result:", rob([]int{1, 2, 3, 1}))
	fmt.Println("result:", rob([]int{2, 7, 9, 3, 1}))
	fmt.Println("result:", rob([]int{1, 2}))

}
