package main
//link :https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
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
func deleteDuplicates(head *ListNode) *ListNode {
   cur :=head
   for {
        if cur == nil {
        	break
		}
		next := cur.Next
		if next !=nil {
			if cur.Val == next.Val {
				cur = next.Next
			}else {
				fmt.Println("value:",cur.Val)
			}
		}
		cur = next
   }

	return head
}
func main() {
  l1 := buildList([]int{1,2,2,3,3,5})
	deleteDuplicates(l1)
}