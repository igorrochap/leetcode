package main

import (
	"fmt"
)

func repeatedSubstringPattern(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if length%(i+1) == 0 {
			appendingTimes := length / (i + 1)
			temp := s[0 : i+1]
			for appendingTimes > 1 {
				temp += s[0 : i+1]
				appendingTimes--
			}
			if temp == s {
				return true
			}
		}
	}
	return false
}

func main() {
	testcases := []string{
		"abab",
		"aba",
		"abcabcabcabc",
	}

	for _, testcase := range testcases {
		fmt.Println(repeatedSubstringPattern(testcase))
	}
}
