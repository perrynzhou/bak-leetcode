143. 重排链表
     给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
     将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.



解题思路

   1.把链表转换为数组tmp
   2.设定l,r :=0,len(tmp)
   3.然后进入for循环，执行如下
​	for l<r	{
​		 tmp[l].Next = tmp[r]
​      	 tmp[r].Next = tmp[l+1]
​      	 l++
​     	 r--
​	}
 	if tmp[l]!=nil {
​		tmp[l].Next = nil 

​	}