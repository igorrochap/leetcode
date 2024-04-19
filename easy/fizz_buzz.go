package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var answer []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				answer = append(answer, "FizzBuzz")
				continue
			}
			answer = append(answer, "Fizz")
			continue
		}
		if i%5 == 0 {
			answer = append(answer, "Buzz")
			continue
		}
		answer = append(answer, strconv.Itoa(i))
	}
	return answer
}

func main() {
	inputs := []int{3, 5, 15}
	for _, input := range inputs {
		fmt.Println(fizzBuzz(input))
	}
}
