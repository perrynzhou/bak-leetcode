package main

//link:https://leetcode-cn.com/problems/palindrome-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	result := true
	nums := make([]int, 0)
	for {
		if head == nil {
			break
		}
		nums = append(nums, head.Val)
		head = head.Next
	}
	l, r := 0, len(nums)-1
	for {
		if l >= r {
			break
		}
		if nums[l] != nums[r] {
			result = false
			break
		}
	}
	return result
}
