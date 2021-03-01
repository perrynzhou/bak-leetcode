package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//[[1,4,5],[1,3,4],[2,6]]
func buildList(nums []int) *ListNode {
	head := &ListNode{}
	tail := head
	for i := 0; i < len(nums); i++ {
		newNode := &ListNode{}
		newNode.Val = nums[i]
		tail.Next = newNode
		tail = newNode
	}
	return head.Next
}
func printLists(l *ListNode) {
	for {
		if l == nil {
			fmt.Println(" ")
			break
		}
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
}

//link:https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	if head != nil && newHead == nil {
		next := head.Next
		newHead = head
		newHead.Next = nil
		head = next
	}
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}
func main() {
	l1 := buildList([]int{1, 2, 3, 4, 5})
	printLists(reverseList(l1))
}
