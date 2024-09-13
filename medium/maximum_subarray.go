package main

import "fmt"

func maxSubArray(nums []int) int {
	localMax := 0
	globalMax := nums[0]
	for _, num := range nums {
		localMax = max(num, localMax+num)
		globalMax = max(globalMax, localMax)
	}
	return globalMax
}

func main() {
	testcases := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{5, 4, -1, 7, 8},
	}

	for _, testcase := range testcases {
		fmt.Printf("%d\n", maxSubArray(testcase))
	}
}
