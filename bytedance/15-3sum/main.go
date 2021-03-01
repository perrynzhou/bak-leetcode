package main

// link:https://leetcode-cn.com/problems/3sum/

import (
	"fmt"
	"sort"
)

func threeSum_v2(nums []int) [][]int {
	result := make([][]int, 0)
	if nums == nil || len(nums) < 3 {
		return result
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println("sort nums:", nums)
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			sum := n1 + n2 + n3
			if sum == 0 {
				result = append(result, []int{n1, n2, n3})
				for l < r && n2 == nums[l] {
					l++
				}
				for l < r && n3 == nums[r] {
					r--
				}
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return result
}
func threeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n < 3 {
		return result
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < n-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 1 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			sum := n1 + n2 + n3
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				result = append(result, []int{n1, n2, n3})
				for l < r && n2 == nums[l] {
					l++
				}
				for l < r && n3 == nums[r] {
					r--
				}
			}
		}
	}
	return result
}
func main() {
	fmt.Println("output result:", threeSum([]int{0, 0, 0, 0}))
	fmt.Println("output result:", threeSum([]int{-2, 0, 1, 1, 2}))
	//[3,0,-2,-1,1,2]
	fmt.Println("output result:", threeSum([]int{3, 0, -2, -1, 1, 2}))
	fmt.Println("output result:", threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))

	//[-2,0,1,1,2]
}
