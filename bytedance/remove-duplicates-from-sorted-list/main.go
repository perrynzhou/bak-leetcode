package main

// link: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var newHead *ListNode
	var tailHead *ListNode
	var next *ListNode

	for {
		if head == nil {
			break
		}
		if head != nil {
			next = head.Next
		}
		head.Next = nil
		if newHead == nil {
			newHead = head
			tailHead = head
		} else {
			if tailHead.Val != head.Val {
				tailHead.Next = head
				tailHead = head
			}
		}
		head = next
	}
	return newHead
}
