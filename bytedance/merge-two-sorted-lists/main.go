package main

// link:https://leetcode-cn.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	if l1 == nil && l2 == nil {
		return nil
	}
	for {
		if l1 == nil && l2 == nil {
			break
		}
		newNode := &ListNode{}
		tail.Next = newNode
		tail = newNode
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				newNode.Val = l1.Val
				l1 = l1.Next
			} else {
				newNode.Val = l2.Val
				l2 = l2.Next
			}
		} else if l1 == nil && l2 != nil {
			newNode.Val = l2.Val
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			newNode.Val = l1.Val
			l1 = l1.Next
		}
	}
	return head.Next
}
