package main

import "strings"

func canConstruct(ransomNote string, magazine string) bool {
	if ransomNote == magazine {
		return true
	}
	for _, letter := range ransomNote {
		if !strings.Contains(magazine, string(letter)) {
			return false
		}
		println("ats -> ", magazine)
		magazine = strings.Replace(magazine, string(letter), "", 1)
		ransomNote = strings.Replace(ransomNote, string(letter), "", 1)
		println("dps -> ", magazine)
		if magazine == "" {
			if ransomNote != "" {
				return false
			}
		}
	}
	return true
}

func main() {
	// ra1, mg1 := "a", "b"
	// ra2, mg2 := "aa", "ab"
	// ra3, mg3 := "aa", "aab"

	// println(canConstruct(ra1, mg1))
	// println(canConstruct(ra2, mg2))
	// println(canConstruct(ra3, mg3))
	// println(canConstruct("aa", "aa"))
	println(canConstruct("aab", "baa"))
	println(canConstruct("abc", "ab"))

	// teste1 := "this is the 1 test"
	// loop := "hello world"

	// for _, char := range loop {
	// 	println(strings.Contains(teste1, string(char)))
	// }
}
