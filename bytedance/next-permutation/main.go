package main

import "fmt"

//link:https://leetcode-cn.com/problems/next-permutation/
/*

解题步骤：
	1.n=len(nums),从右边l=len(nums)-2,n=l+1开始比较，如果nums[l]>=nums[l+1],需要l--
	2.如果l为-1，说明是最大的序列，直接进入第3步
	3.如果l>=0, r:=n-1,如果nums[l]>=nums[r],r--,进行交换nums[l]和nums[r]
	4.start:=l+1,end:=n-1,进行数组反转

*/

// 5,4,6,6,1
//5,6,1,4,6
func nextPermutation(nums []int) {
	// 1,2,2->2,2,1->2,1,2
	// 5,4,6,6,1 -> 5,6,6,4,1->5,6,1,4,6
	n := len(nums)
	l := n - 2
	//从右边开始比较两个元素的大小，nums[l]>=nums[l+1],前一个元素大于后一个l--
	for l >= 0 && nums[l] >= nums[l+1] {
		l--
	}
	fmt.Println("origin:", nums)

	// l=1,nums[l]=4
	if l >= 0 {
		// nums[r] = =4
		//	原始 5,4,6,6,1
		// 从右往左找，找到了nums[l]<nums[l+1]的元素，记录下前一个元素的小标
		r := n - 1
		// 这个元素和从最右边比较，再次查找比nums[l]小的元素，往左边移动r的位置
		for r >= 0 && nums[l] >= nums[r] {
			r--
		}
		//r=3
		//5,6,6,4,1
		//然后在交换nums[l]和nums[r]的位置
		nums[l], nums[r] = nums[r], nums[l]

	}
	// 如果所有的元素都是前一个比后一个大，则是最大序列，l=-1,需要返回整个数组
	//5,6,1,4,6
	start, end := l+1, n-1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	fmt.Println("result:", nums)
}
func main() {
	//	nextPermutation([]int{5,3,2,1})

	//nextPermutation([]int{1,1,2,3})
	//nextPermutation([]int{1,3,2})
	//nextPermutation([]int{1,2,2})
	nextPermutation([]int{3, 2, 1})
	nextPermutation([]int{5, 4, 6, 6, 1})

}
