package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	arrayLength := len(nums)
	for i, element := range nums {
		if element == val {
			nums[i] = 999
			arrayLength--
			continue
		}
	}
	sort.Ints(nums)
	return arrayLength
}

type Tescase struct {
	Value    int
	Elements []int
}

func main() {
	testcases := []Tescase{
		Tescase{Value: 3, Elements: []int{3, 2, 2, 3}},
		Tescase{Value: 2, Elements: []int{0, 1, 2, 2, 3, 0, 4, 2}},
	}

	for _, testcase := range testcases {
		output := removeElement(testcase.Elements, testcase.Value)
		fmt.Println(output)
		fmt.Println(testcase.Elements)
	}
}
