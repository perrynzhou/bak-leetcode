package main

// link:https://leetcode-cn.com/problems/merge-k-sorted-lists/
import (
	"fmt"
	"math"
)

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
func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	tail := head
	var emptyLen int

	for {
		fmt.Println("###emptyLen=:", emptyLen, "listLen:", len(lists))

		min := math.MaxInt32
		var index int
		emptyLen := 0
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				if min >= lists[i].Val {
					min, index = lists[i].Val, i
					fmt.Println("-debug---min:", min, ",i:", i)
				}
			} else {
				emptyLen++
			}
		}
		if emptyLen >= len(lists) {
			break
		}
		newNode := lists[index]
		fmt.Println("min:", min, ",emptyLen:", emptyLen, ",index:", index, ",listLen:", len(lists))
		if lists[index] != nil {
			fmt.Printf("lists[%d]:%v\n", index, lists[index])
			lists[index] = lists[index].Next
			fmt.Printf("next lists[%d]:%v\n", index, lists[index])

		}

		newNode.Next = nil
		fmt.Println("index:", index, ",min:", min)
		if tail == nil {
			tail = newNode
			head = tail
		} else {
			tail.Next = newNode
			tail = newNode
		}

	}
	return head
}

///[[1,4,5],[1,3,4],[2,6]]
func main() {

	l1 := buildList([]int{1, 4, 5})
	l2 := buildList([]int{1, 3, 4})

	l3 := buildList([]int{2, 6})
	printLists(l1)
	printLists(l2)
	printLists(l3)
	fmt.Println("-------------")
	lists := []*ListNode{l1, l2, l3}
	l6 := mergeKLists(lists)
	printLists(l6)

}
