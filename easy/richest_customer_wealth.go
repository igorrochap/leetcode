package main

func maximumWealth(accounts [][]int) int {
	var maxWealth int
	for i, row := range accounts {
		currentWealth := 0
		for _, value := range row {
			currentWealth += value
		}
		if i == 0 {
			maxWealth = currentWealth
			continue
		}
		if currentWealth > maxWealth {
			maxWealth = currentWealth
		}
	}
	return maxWealth
}

func main() {
	accounts1 := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	accounts2 := [][]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}
	accounts3 := [][]int{
		{2, 8, 7},
		{7, 1, 3},
		{1, 9, 5},
	}

	println(maximumWealth(accounts1))
	println(maximumWealth(accounts2))
	println(maximumWealth(accounts3))
}
