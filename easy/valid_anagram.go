package main

func isAnagram(s string, t string) bool {
	letterCount := make(map[rune]int)
	validAnagram := true

	if len(s) != len(t) {
		return false
	}

	for _, letter := range s {
		_, inMap := letterCount[letter]
		if inMap {
			letterCount[letter]++
			continue
		}
		letterCount[letter] = 1
	}
	for _, letter := range t {
		_, inMap := letterCount[letter]
		if !inMap {
			validAnagram = false
			break
		}
		if letterCount[letter] == 0 {
			validAnagram = false
			break
		}
		letterCount[letter]--
	}
	return validAnagram
}

func main() {
	s1, t1 := "anagram", "nagaram"
	s2, t2 := "rat", "car"
	s3, t3 := "ab", "a"
	println(isAnagram(s1, t1))
	println(isAnagram(s2, t2))
	println(isAnagram(s3, t3))
}
