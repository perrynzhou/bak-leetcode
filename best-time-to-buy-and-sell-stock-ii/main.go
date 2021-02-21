package main

//link:https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	best := 0
	for i := 1; i < n-1; i++ {
		if prices[i] > prices[i-1] {
			best = best + prices[i] - prices[i-1]
		}
	}
	return best
}
