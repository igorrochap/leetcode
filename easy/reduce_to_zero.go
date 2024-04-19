package main

func numberOfSteps(num int) int {
	steps := 0
	for {
		if num == 0 {
			break
		}

		steps++
		if num%2 == 0 {
			num /= 2
			continue
		}
		num -= 1
	}
	return steps
}

func main() {
	number1 := 14
	number2 := 8
	number3 := 123

	println(numberOfSteps(number1))
	println(numberOfSteps(number2))
	println(numberOfSteps(number3))
}
