package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
    s:=make([]int,0)
    for i:=0;i<len(nums1);i++ {
    	s= append(s,nums1[i])
	}
	for i:=0;i<len(nums2);i++ {
		s= append(s,nums2[i])
	}
	sort.Slice(s,func(i,j int)bool{
		return s[i]>s[j]
	})
	if len(s)%2 == 0 {
		result = (float64(s[len(s)/2-1])+float64(s[len(s)/2]))/2
	}else {
		result= float64(s[len(s)/2])
	}
	return result
}
