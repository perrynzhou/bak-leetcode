package main

import "fmt"

//link: https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
func maxSubArray(nums []int) int {
	sum,max :=0,nums[0]
	for i:=1;i<len(nums);i++ {
		if sum+nums[i] <nums[i] {
			sum = nums[i]
		}else {
			sum +=nums[i]
			if sum>max {
				max = sum
			}
		}
	}
	return max
}
func main() {
	nums :=[]int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println("max:",maxSubArray(nums))
}