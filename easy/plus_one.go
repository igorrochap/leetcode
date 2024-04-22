package main

import "fmt"

func plusOne(digits []int) []int {
	length := len(digits)
	digits[length-1] += 1
	if digits[length-1] == 10 {
		for i := length - 1; i >= 0; i-- {
			if i == 0 && digits[i] == 10 {
				digits = append(digits, 0)
				digits[0] = 1
			}
			if digits[i] == 10 {
				digits[i] = 0
				digits[i-1] += 1
			}
		}
	}
	return digits
}

func main() {
	testcases := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
		{1, 2, 9},
		{2, 9, 9},
		{9, 9, 9},
	}

	for _, testcase := range testcases {
		fmt.Println(plusOne(testcase))
	}
}
