package main
import (
	"fmt"
)
// link:https://leetcode-cn.com/problems/reorder-list/


type ListNode struct {
	Val  int
	Next *ListNode
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
func reorderList_v2(head *ListNode)  {
   tmp :=[]*ListNode{}
   cur :=head
   for {
   	 if cur ==nil {
   	 	break
	 }
	   next := cur.Next
	   cur.Next = nil
	   tmp = append(tmp,cur)
	  cur = next

   }

   j := len(tmp)-1
   first := tmp[:len(tmp)/2]
   second :=tmp[len(tmp)/2:]
   for i:=len(tmp)/2;i<=j;i++ {
   	     if j>=0 {
   	     	tmp[i],tmp[j] = tmp[j],tmp[i]
   	     	j--
		 }
   }

   for i:=0;i<len(first);i++ {
   	  fmt.Println(first[i].Val)
   }
   fmt.Println("--------")
	for i:=0;i<len(second);i++ {
		fmt.Println(second[i].Val)
	}
	j =0
	var tail *ListNode
	for i:=0;i<len(first);i++ {
		if tail == nil {
			tail = first[i]
			if j<len(second) {
				tail.Next = second[j]
				tail = second[j]
				j++
			}
			continue
		}
		tail.Next = first[i]
		tail = first[i]
		if j<len(second) {
			tail.Next = second[j]
			tail = second[j]
			j++
		}
	}
	if tail !=nil && j <len(second) {
		tail.Next = second[j]
	}
}
func reorderList(head *ListNode) {
	list :=[]*ListNode{}
	cur :=head
	for cur!=nil {
		list = append(list,cur)
		cur=cur.Next
	}
	if len(list) ==0 {
		return
	}
	l,r := 0,len(list)-1
	for l<r {
		list[l].Next =list[r]
		list[r].Next = list[l+1]
		l++
		r--
	}
	if list[l]!=nil {
		list[l].Next = nil
	}
}
func main() {
	l1:=BuildList([]int{1,2,3,4})
	reorderList(l1)
	PrintList(l1)
}
