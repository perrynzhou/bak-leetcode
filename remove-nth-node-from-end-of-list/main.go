package main

// link https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var prev *ListNode
	var next *ListNode
	result := head
	var index int
	var count int
	for {
		if head == nil {
			break
		}
		count++
		head = head.Next
	}
	if count == 1 && n == 1 {
		return nil
	}

	if count > 1 && count == n {
		result = result.Next
		return result
	}
	head = result
	for {
		if head == nil {
			break
		}
		index++
		if count-index != n-1 {
			prev = head
		}
		if count-index == n-1 {
			next = head.Next
			break
		}
		head = head.Next
	}
	if prev != nil {
		prev.Next = next
	}
	return result
}
