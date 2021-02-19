package main

import "sort"

//link:https://leetcode-cn.com/problems/sort-list/

  type ListNode struct {
      Val int
      Next *ListNode
  }

func sortList(head *ListNode) *ListNode {

	tmp := head
	nums := make([]int,0)
	for {
		if tmp == nil {
			break
		}
		nums = append(nums,tmp.Val)
		next := tmp.Next
		tmp = next
	}
	sort.Slice(nums,func(i,j int)bool {
		return nums[i]<nums[j]
	})
	tmp = head
	for i:=0;i<len(nums);i++ {
		tmp.Val = nums[i]
		tmp = tmp.Next
	}
	return head
}