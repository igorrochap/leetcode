package main

func majorityElement(nums []int) int {
	elementOcurrences := make(map[int]int)
	for _, number := range nums {
		if _, exists := elementOcurrences[number]; !exists {
			elementOcurrences[number] = 1
			continue
		}
		elementOcurrences[number] += 1
	}
	for element, ocurrences := range elementOcurrences {
		if ocurrences > (len(nums) / 2) {
			return element
		}
	}
	return -1
}

func main() {
	testcases := [][]int{
		{3, 2, 3},
		{2, 2, 1, 1, 1, 2, 2},
	}

	for _, testcase := range testcases {
		println(majorityElement(testcase))
	}
}
