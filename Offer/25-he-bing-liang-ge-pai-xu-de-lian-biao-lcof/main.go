package main

import (
	"fmt"
	"sort"
)

//link: https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/

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

func mergeTwoLists_v1(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := make([]*ListNode, 0)
	for l1 != nil {
		tmp = append(tmp, l1)
		l1 = l1.Next
	}
	for l2 != nil {
		tmp = append(tmp, l2)
		l2 = l2.Next
	}
	if len(tmp) == 0 {
		return nil
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].Val < tmp[j].Val
	})

	newHead := tmp[0]
	newTail := newHead
	for i := 1; i < len(tmp); i++ {
		newTail.Next = tmp[i]
		newTail = tmp[i]
	}
	if newTail != nil {
		newTail.Next = nil
	}
	return newHead
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var newHead *ListNode
	var newTail *ListNode

	for {
		if l1==nil && l2 ==nil {
			break
		}
		var newNode *ListNode
		if l1!=nil && l2 != nil {
			if l1.Val < l2.Val  {
				newNode = l1
				l1 = l1.Next
			}else{
				newNode = l2
				l2 = l2.Next
			}
		}else if l1!=nil && l2 ==nil {
			newNode = l1
			l1 = l1.Next
		}else if l1==nil&&l2 !=nil {
			newNode = l2
			l2 = l2.Next
		}
		if newHead == nil {
			newHead = newNode
			newTail = newNode
		}else {
			newTail.Next = newNode
			newTail = newNode
		}
	}
	return newHead
}
func main() {
	l1 := buildList([]int{1, 2, 4})
	l2 := buildList([]int{1, 3, 4})
	newHead := mergeTwoLists(l1, l2)
	printLists(newHead)
}
