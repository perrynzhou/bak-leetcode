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
   dict :=make(map[int]int)
   for {
        if cur == nil {
        	break
		}
		dict[cur.Val]++
		cur = cur.Next
   }
   var newHead *ListNode
   var newTail *ListNode
   cur = head
    for {
    	if cur ==nil {
    		break
		}
		next := cur.Next
		cur.Next = nil
		if dict[cur.Val] ==1 {
			if newHead == nil {
				newHead = cur
				newTail = cur
				fmt.Println("head:",cur.Val)
			}else {
				newTail.Next = cur
				newTail = cur
				fmt.Println("tail:",cur.Val)
			}
		}
		cur = next
	}
	return newHead
}
func deleteDuplicates_v2(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for {
		if cur == nil {
			break
		}
		next := cur.Next
		if prev == nil {
			prev = cur
		}else {
			if prev.Val == cur.Val {
				cur = next
				prev = nil
			}else {
				fmt.Println()
			}
		}

	}
}
func main() {
  l1 := buildList([]int{1,2,2,3,3,5})
  l2:=deleteDuplicates(l1)
	printLists(l2)
}