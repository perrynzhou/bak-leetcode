package main

import "fmt"

// link:https://leetcode-cn.com/problems/swap-nodes-in-pairs/

type ListNode struct {
	Val  int
	Next *ListNode
}

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

//head = [1,2,3,4]
//输出：[2,1,4,3]
func swapPairs(head *ListNode) *ListNode {
	var newHead *ListNode
	var first *ListNode
	var second *ListNode
	if head == nil {

		return nil
	}
	if head.Next == nil {
		return head
	}
	tail := newHead
	for {
		if head == nil {
			break
		}
		if first == nil {
			first = head
		} else {
			second = head
		}
		head = head.Next
		if first != nil && second != nil {
			if tail == nil {
				tail = second
				newHead = tail
			} else {
				tail.Next = second
				tail = tail.Next
			}
			tail.Next = first
			tail = tail.Next
			tail.Next = nil
			fmt.Printf("first:%d,", first.Val)
			fmt.Printf("second:%d\n", second.Val)
			first = nil
			second = nil
		}
	}
	if first != nil {
		tail.Next = first
		tail = tail.Next
		tail.Next = nil
		fmt.Printf("last:%d\n", first.Val)

	}
	return newHead
}
func main() {
	l1 := buildList([]int{1, 2, 3, 4, 5})
	printLists(l1)
	l2 := swapPairs(l1)
	printLists(l2)
}
