package main
import (
	"fmt"
)
//link: https://leetcode-cn.com/problems/binary-search/
func search(nums []int, target int) int {
   l,r :=0,len(nums)-1
   index :=-1
   for l<=r {
   	  mid :=(l+r)/2
   	  if target == nums[mid] {
   	  	index = mid
   	  	break
	  }else if target > nums[mid] {
	  	l = mid+1
	  }else {
	  	r = mid-1
	  }
   }
   return index
}
func main() {
	fmt.Println("index:",search([]int{-1,0,3,5,9,12},9))
}