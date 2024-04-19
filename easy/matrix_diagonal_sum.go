package main

func diagonalSum(mat [][]int) int {
	sum := 0
	matrixLength := len(mat)
	for i := range mat {
		sum += mat[i][i]
		matrixLength--
		if len(mat)%2 != 0 && (i == matrixLength) {
			continue
		}
		sum += mat[i][matrixLength]
	}
	return sum
}

func main() {
	mat1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	mat2 := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
	mat3 := [][]int{
		{5},
	}

	println(diagonalSum(mat1))
	println(diagonalSum(mat2))
	println(diagonalSum(mat3))
}
