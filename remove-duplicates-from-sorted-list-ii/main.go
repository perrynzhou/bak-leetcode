package main
//link :https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
   dict :=make(map[int]int)
   tmp :=head
   for {
   	   if tmp ==nil {
   	    	break
	   }
	   dict[tmp.Val]++
	   tmp = tmp.Next
   }
   tmp = head
   var prev *ListNode
   for k,v :=range dict {
   	if v==1 {
   		tmp.Val=k
   		prev = tmp
   		tmp = tmp.Next
	}
   }
   if prev!=nil {
	   prev.Next = nil
   }
	return head
}