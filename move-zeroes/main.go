package main

import "fmt"

//link :https://leetcode-cn.com/problems/move-zeroes/

func moveZeroes(nums []int) {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for ; j < n; j++ {
		nums[j] = 0
	}
	fmt.Println("nums:", nums)

}
func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
