package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	testcases := map[string]string{
		"sadbutsad": "sad",
		"leetcode":  "leeto",
	}

	for haystack, needle := range testcases {
		fmt.Println(strStr(haystack, needle))
	}
}
