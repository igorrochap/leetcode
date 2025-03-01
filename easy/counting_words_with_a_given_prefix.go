package main

import (
	"fmt"
)

type Testcase struct {
	Prefix string
	Words  []string
}

type TrieNode struct {
	children map[rune]*TrieNode
	counter  int
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
		node.counter++
	}

}

func (t *Trie) OccurrencesOf(prefix string) int {
	node := t.root
	for _, char := range prefix {
		if node.children[char] == nil {
			return 0
		}
		node = node.children[char]
	}
	return node.counter
}

func prefixCount(words []string, pref string) int {
	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}
	return trie.OccurrencesOf(pref)
}

func main() {
	testcases := []Testcase{
		{Prefix: "at", Words: []string{"pay", "attention", "practice", "attend"}},
		{Prefix: "code", Words: []string{"leetcode", "win", "loops", "success"}},
	}
	for _, testcase := range testcases {
		fmt.Println(prefixCount(testcase.Words, testcase.Prefix))
	}
}
