package main

import (
	"math"
	"sort"
)

// link https://leetcode-cn.com/problems/3sum-closest/

func threeSumClosest_v2(nums []int, target int) int {
	best := math.MaxInt32
	if nums == nil || len(nums) < 3 {
		return best
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			sum := n1 + n2 + n3
			if math.Abs(float64(sum-target)) < math.Abs(float64(best-target)) {
				best = sum
			}
			dis := sum - target
			if dis == 0 {
				return sum
			} else if dis > 0 {
				for l < r && n3 == nums[r] {
					r--
				}
			} else {
				for l < r && n2 == nums[l] {
					l++
				}
			}
		}
	}
	return best
}
func threeSumClosest(nums []int, target int) int {
	best,n := math.MaxInt32,len(nums)
	 if n<3 {
	 	return  best
	 }
	 sort.Slice(nums,func(i,j int)bool{
	 	return nums[i]<nums[j]
	 })
	for i:=0;i<n-2;i++ {
		n1 := nums[i]
		if i>1 && n1==nums[i-1] {
			continue
		}
		l,r := i+1,n-1
		for l<r {
			n2,n3 := nums[l],nums[r]
			sum := n1+n2+n3
			if math.Abs(float64(sum-target)) < math.Abs(float64(best-target)) {
				best = sum
			}
			dis := sum - target
			if dis > 0 {
				for l<r && n3 ==nums[r] {
					r--
				}
			}else if dis < 0 {
				for l<r && n2 ==nums[l] {
					l++
				}
			}else {
				return sum
			}
		}
	}
	return best
}