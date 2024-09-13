package main

import "fmt"

func isSameAfterReversals(num int) bool {
	return num%10 != 0 || num == 0
}

func main() {
	testcases := []int{526, 1800, 0}
	for _, test := range testcases {
		fmt.Println(isSameAfterReversals(test))
	}
}
