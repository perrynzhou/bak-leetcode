package main

import "fmt"

//link:https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

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
func getKthFromEnd(head *ListNode, k int) *ListNode {
	var newHead *ListNode
	cur := head
	count, i := 0, 0
	for cur != nil {
		cur = cur.Next
		count++
	}
	for head != nil {
		if (count - i) == k {
			newHead = head
			break
		}
		head = head.Next
		i++
	}
	return newHead

}
func main() {
	l1 := buildList([]int{1, 2, 3, 4, 5})
	l2 := getKthFromEnd(l1, 3)
	printLists(l2)
}
