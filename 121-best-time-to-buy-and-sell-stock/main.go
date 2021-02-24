package main

import "fmt"

//link:https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit_v2(prices []int) int {
	n := len(prices)
	res, min := 0, prices[0]
	for i := 1; i < n; i++ {
		if tmp := prices[i] - min; tmp > res {
			res = tmp
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return res
}

func maxProfit(prices []int) int {
	min, result, n := prices[0], 0, len(prices)
	for i := 1; i < n; i++ {
		if tmp := prices[i] - min; tmp > result {
			result = tmp
		}
		if prices[i] > min {
			min = prices[i]
		}
	}
	return result
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{2, 4, 1}))

}
