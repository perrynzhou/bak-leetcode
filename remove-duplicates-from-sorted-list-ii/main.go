package main
//link :https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	originHead := head
   slow,fast := head,head.Next
   var newHead *ListNode
	var newTail *ListNode
 // 1,1,2,2,3
	for slow != nil && fast !=nil {
    	if slow.Val != fast.Val {
    		if newHead == nil {
    			newHead = slow
				newTail = slow
			}else{
				newTail.Next = slow
				newTail=slow
			}
		}else {
			slow = fast.Next

		}
		slow = slow.Next
		fast = fast.Next
	}
}