package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	buy := prices[0]
	for _, price := range prices {
		if price < buy {
			buy = price
			continue
		}
		localProfit := price - buy
		if localProfit > profit {
			profit = localProfit
		}
	}
	return profit
}

func main() {
	testcases := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
	}

	for _, testcase := range testcases {
		fmt.Printf("%d\n", maxProfit(testcase))
	}
}
