package main

//link: https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
import (
	"fmt"
)

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
func reverseKGroup(head *ListNode, k int) *ListNode {
	var newHead *ListNode
	var newTail *ListNode
	var tmpHead *ListNode
	var tmpTail *ListNode
	var next *ListNode
	var count int
	cur := head
	for {
		if cur == nil {
			break
		}
		count++
		cur = cur.Next
	}
	round := count / k
	for i := 0; i < round; i++ {
		if head == nil {
			break
		}
		tmpHead = nil
		tmpTail = nil
		for j := 0; j < k; j++ {
			next = head.Next
			newNode := head
			newNode.Next = nil
			if tmpHead == nil {
				tmpHead = head
				tmpTail = tmpHead
			} else {
				newNode.Next = tmpHead
				tmpHead = newNode
			}
			head = next
		}
		if newHead == nil {
			newHead = tmpHead
			newTail = tmpTail
		} else {
			newTail.Next = tmpHead
			newTail = tmpTail

		}
		fmt.Println("newHead:", newHead.Val, ",newTail:", newTail.Val)

		//fmt.Println("tmpHead:",tmpHead.Val,",tmpTail:",tmpTail.Val)
		printLists(newHead)
	}
	if next != nil {
		newTail.Next = next
		fmt.Println("next:", next.Val)
		printLists(newHead)
		//fmt.Println("newHead:",newHead.Val)

	}
	return newHead
}
func main() {
	l1 := buildList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	printLists(l1)
	reverseKGroup(l1, 4)
	//printLists(l2)
}
