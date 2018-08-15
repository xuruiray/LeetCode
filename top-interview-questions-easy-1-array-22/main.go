package main

import "fmt"

func main() {
	// prices := []int{7, 1, 5, 3, 6, 4, 9, 10}
	prices := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(prices)
	fmt.Println(result)
}

func maxProfit(prices []int) int {
	sum := 0
	for i, v := range prices {
		if i == 0 {
			continue
		}
		if prices[i]-prices[i-1] > 0 {
			sum += v - prices[i-1]
		}
	}
	return sum
}
