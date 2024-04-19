package main

import (
	"sort"
)

func mostFrequentEven(nums []int) int {
	elementOcurrences := make(map[int]int)
	for _, number := range nums {
		if number%2 == 0 {
			if _, exists := elementOcurrences[number]; !exists {
				elementOcurrences[number] = 1
				continue
			}
			elementOcurrences[number] += 1
		}
	}

	if len(elementOcurrences) > 0 {
		return getElement(elementOcurrences)
	}

	return -1
}

func getElement(elementOcurrences map[int]int) int {
	var keys []int
	for k := range elementOcurrences {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	sort.SliceStable(keys, func(i, j int) bool {
		return elementOcurrences[keys[i]] > elementOcurrences[keys[j]]
	})

	return keys[0]
}

func main() {
	testcases := [][]int{
		{0, 1, 2, 2, 4, 4, 1},
		{4, 4, 4, 9, 2, 4},
		{29, 47, 21, 41, 13, 37, 25, 7},
		{8154, 9139, 8194, 3346, 5450, 9190, 133, 8239, 4606, 8671, 8412, 6290},
	}

	for _, testcase := range testcases {
		println(mostFrequentEven(testcase))
	}
}
