package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	n := len(nums)
	max := math.MinInt32
	if n == 1 {
		return nums[n-1]
	}
	for i := 0; i < n; i++ {
		tmp := nums[i]
		if tmp > max {
			max = tmp
		}
		for j := i + 1; j < n; j++ {
			tmp += nums[j]
			if tmp > max {
				max = tmp
			}
		}
	}
	fmt.Println(" end max:", max)
	return max
}
func maxSubArray_v2(nums []int) int {

	n := len(nums)
	if n == 1 {
		return nums[n-1]
	}
	max := math.MinInt32
	for i := 0; i < n; i++ {
		tmp := nums[i]
		if tmp > max {
			max = tmp
		}
		for j := i + 1; j < n; j++ {
			tmp += nums[j]
			if tmp > max {
				max = tmp
			}
		}
	}
	fmt.Println("max:", max)
	return max
}
func main() {
	/*
		maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
		maxSubArray([]int{-2, 1})
		maxSubArray([]int{-1, 0, -2})


	*/
	//[-2,1,-3,4,-1,2,1,-5,4]
	maxSubArray_v2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	maxSubArray_v2([]int{-2, 1})
	maxSubArray_v2([]int{-1, 0, -2})
}
