package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//link:https://leetcode-cn.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	dict := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := dict[head]; ok {
			return true
		}
		dict[head] = struct{}{}
		head = head.Next
	}
	return false
}
func hasCycle_v2(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
