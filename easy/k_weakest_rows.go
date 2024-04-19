package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func kWeakestRows(mat [][]int, k int) []int {
	rowsSoldierQuantity := make(PairList, len(mat))
	var kWeakestIndexes []int
	for i, row := range mat {
		soldierQuantity := 0
		for _, person := range row {
			if person == 1 {
				soldierQuantity += 1
			}
		}
		rowsSoldierQuantity[i] = Pair{i, soldierQuantity}
	}

	sort.Slice(rowsSoldierQuantity, func(i, j int) bool {
		if rowsSoldierQuantity[i].Value != rowsSoldierQuantity[j].Value {
			return rowsSoldierQuantity[i].Value < rowsSoldierQuantity[j].Value
		}
		return rowsSoldierQuantity[i].Key < rowsSoldierQuantity[j].Key
	})

	for _, row := range rowsSoldierQuantity {
		if len(kWeakestIndexes) == k {
			break
		}
		kWeakestIndexes = append(kWeakestIndexes, row.Key)
	}

	return kWeakestIndexes
}

func main() {
	row1 := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}

	fmt.Println(kWeakestRows(row1, 3))
}
