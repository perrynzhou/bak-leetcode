package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers_v2(l1 *ListNode, l2 *ListNode) *ListNode {
	mid := 0
	var tail *ListNode
	var head *ListNode
	// 题目的结果是按照尾插发来插入链表的
	// 遍历l2和l2,使用中间变量mid来进位
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			mid += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			mid += l2.Val
			l2 = l2.Next
		}
		// 如果2个值进行相加，大于10，则链表的头的值为mid%10的值
		value := mid % 10
		if tail == nil {
			tail = &ListNode{
				Val:  value,
				Next: nil,
			}
			head = tail
		} else {
			tail.Next = &ListNode{
				Val:  value,
				Next: nil,
			}
			tail = tail.Next
		}
		// 当插入完成以后，需要针对这个值进行进位，进位的值为mid/10；如果mid小于10，则设置mid为0
		if mid >= 10 {
			mid = mid / 10
		} else {
			mid = 0
		}

	}
	//最后需要针对mid判断，如果大于0，则表示还有一个节点需要插入
	if mid > 0 {
		tail.Next = &ListNode{
			Val:  mid,
			Next: nil,
		}
		tail = tail.Next
	}
	return head
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	var sum int
	var incr int
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += incr
		if sum > 9 {
			incr = 1
		} else {
			incr = 0
		}
		newNode := &ListNode{
			Val: sum % 10,
		}
		fmt.Println("mid:", newNode.Val, ",sum:", sum, ",incr:", incr)
		if head == nil {
			head = newNode
			tail = head
		} else {
			tail.Next = newNode
			tail = newNode
		}
		sum = 0
	}
	if incr == 1 {
		newNode := &ListNode{
			Val: 1,
		}
		tail.Next = newNode
	}
	return head
}
func BuildList(nums []int) *ListNode {
	var tail *ListNode
	var head *ListNode
	for _, v := range nums {
		if head == nil {
			head = &ListNode{
				Val:  v,
				Next: nil,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val:  v,
				Next: nil,
			}
			tail = tail.Next
		}
	}
	return head
}
func PrintList(l1 *ListNode) {
	for {
		if l1 == nil {
			fmt.Println("")
			break
		}
		fmt.Print(l1.Val, " ")
		l1 = l1.Next
	}
}
func main() {
	l1 := BuildList([]int{9, 9, 9})
	l2 := BuildList([]int{9, 9, 9, 9, 9, 9, 9})
	PrintList(l1)
	PrintList(l2)
	head := addTwoNumbers(l1, l2)
	PrintList(head)
}
