package main

import "fmt"

type TrieNode struct {
	children    []TrieNode
	isEndOfWord bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{children: make([]TrieNode, 26), isEndOfWord: false}
}

func insertOnTrie(key string, root *TrieNode) {
	crawlPointer := root
	var index int
	for level := 0; level < len(key); level++ {
		index = int(key[level] - 'a')
		if len(crawlPointer.children[index].children) == 0 {
			crawlPointer.children[index] = *newTrieNode()
		}
		crawlPointer = &crawlPointer.children[index]
	}
	crawlPointer.isEndOfWord = true
}

func countChildren(node TrieNode, index *int) int {
	count := 0
	for i := 0; i < 26; i++ {
		if len(node.children[i].children) > 0 {
			count++
			*index = i
		}
	}
	return count
}

func walkOnTrie(root *TrieNode) string {
	crawlPointer := *root
	var index int
	var prefix string
	for {
		if countChildren(crawlPointer, &index) != 1 || crawlPointer.isEndOfWord == true {
			break
		}
		crawlPointer = crawlPointer.children[index]
		prefix = prefix + (string(rune('a' + index)))
	}
	return prefix
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	root := *newTrieNode()
	for i := 0; i < len(strs); i++ {
		insertOnTrie(strs[i], &root)
	}
	return walkOnTrie(&root)
}

func main() {
	testcases := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"", "b"},
		{"a"},
		{"ab", "a"},
	}

	for _, testcase := range testcases {
		prefix := longestCommonPrefix(testcase)
		fmt.Println(prefix)
	}
}
