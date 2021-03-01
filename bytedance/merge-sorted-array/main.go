package main

import (
	"fmt"
	"sort"
)

//link:https://leetcode-cn.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {

	j := 0
	for i := m; i < len(nums1); i++ {
		if j >= n {
			break
		}
		nums1[i] = nums2[j]
		j++
	}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	fmt.Println("nums1:", nums1)
}
func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
