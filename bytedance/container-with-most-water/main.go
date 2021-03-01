package main

import "fmt"

func maxArea(height []int) int {
	// 定义min为数组左边界和右边界最小值
	var min int
	// 定义max为数组计算面积的最大值，初始化为0
	var max int
	// 从最左开始
	i := 0
	//同时定义右边界的起始位置
	j := len(height) - 1
	for {
		if i == j {
			break
		}
		// 取数据左右最小的值
		if height[i] > height[j] {
			min = height[j]
		} else {
			min = height[i]
		}
		// 计算面积和当前max进行比较
		if (j-i)*min > max {
			max = (j - i) * min
		}
		// 左边界元素比右边界大，需要移动右边界j减1，否则左边界i加1
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return max
}
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
