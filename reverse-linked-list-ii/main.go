package main

import "fmt"

// link: https://leetcode-cn.com/problems/reverse-linked-list-ii/

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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var prev *ListNode
	curHead := head
	i := 1
	//1->2->3->4->5->NULL, m = 1, n = 4
	for {
		if i >= m {
			break
		}
		prev = head
		head = head.Next
		i++
	}

	if prev != nil {
		fmt.Println("prev:", prev.Val)
		head = prev.Next
		prev.Next = nil
	}
	var newHead *ListNode
	var newTail *ListNode
	i, j := 0, n-m
	for {
		if i > j || head == nil {
			break
		}
		newNode := head
		head = head.Next
		newNode.Next = nil
		if newHead == nil {
			newHead = newNode
			newTail = newNode
		} else {
			newNode.Next = newHead
			newHead = newNode
		}
		fmt.Println("newHead:", newHead.Val)
		i++
	}
	if head != nil {
		newTail.Next = head
	}
	if prev != nil {
		prev.Next = newHead
		newHead = curHead
	}
	return newHead
}
func main() {
	l1 := buildList([]int{1, 2, 3, 4, 5})
	l3 := reverseBetween(l1, 2, 5)
	printLists(l3)
}
