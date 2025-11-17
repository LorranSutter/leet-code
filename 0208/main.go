package main

import (
	"fmt"
	"runtime"
)

// The idea here is to use Left-child right-sibling binary tree (LCRS)
// The child of a node belongs to the same word
//   e.g. for "jam", the child of 'a' is 'j', the child of 'j' is 'm', and the child of'm' is 'a'
// The siblings of a node are different words that share the same prefix
//   e.g. for "jam" and "jay", 'm' is a sibling of 'y' and share the same prefix 'ja'

type TrieNode struct {
	Letter      byte
	Next        *TrieNode
	Child       *TrieNode
	IsEndOfWord bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	trie := Trie{
		root: &TrieNode{},
	}
	return trie
}

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	prevNode := t.root
	node := t.root
	for len(word) > 0 {
		prevNode = node

		if node.Letter == 0 {
			node.Letter = word[0]
			node.Child = &TrieNode{}
			node = node.Child
			word = word[1:]
		} else if node.Letter == word[0] {
			if node.Child == nil {
				node.Child = &TrieNode{}
			}
			node = node.Child
			word = word[1:]
		} else if node.Next == nil {
			node.Next = &TrieNode{}
			node = node.Next
		} else {
			node = node.Next
		}
	}

	prevNode.IsEndOfWord = true
}

func (t *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	prevNode := t.root
	node := t.root
	for len(word) > 0 {
		prevNode = node

		if node == nil || node.Letter == 0 {
			return false
		}
		if node.Letter == word[0] {
			node = node.Child
			word = word[1:]
		} else {
			node = node.Next
		}
	}

	return prevNode.IsEndOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}

	node := t.root
	for len(prefix) > 0 {
		if node == nil || node.Letter == 0 {
			return false
		}
		if node.Letter == prefix[0] {
			node = node.Child
			prefix = prefix[1:]
		} else {
			node = node.Next
		}
	}

	return true
}

func main() {
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	obj := Constructor()
	obj.Insert("a")
	obj.Insert("apple")
	obj.Insert("bunda")
	fmt.Println(obj.Search("apple") == true)
	fmt.Println(obj.Search("app") == false)
	fmt.Println(obj.StartsWith("app") == true)
	obj.Insert("app")
	fmt.Println(obj.root)
	fmt.Println(obj.root.Next)
	fmt.Println(obj.root.Next.Next)
	fmt.Println(obj.Search("app") == true)
	fmt.Println(obj.Search("a") == true)
	fmt.Println(obj.StartsWith("a") == true)

	fmt.Println()

	obj = Constructor()
	obj.Insert("hello")
	fmt.Println(obj.Search("hello") == true)
	fmt.Println(obj.Search("helloa") == false)
	fmt.Println(obj.Search("hello") == true)
	fmt.Println(obj.StartsWith("hell") == true)
	fmt.Println(obj.StartsWith("helloa") == false)
	fmt.Println(obj.StartsWith("hello") == true)

	fmt.Println()

	obj = Constructor()
	obj.Insert("app")
	obj.Insert("apple")
	obj.Insert("beer")
	obj.Insert("add")
	obj.Insert("jam")
	obj.Insert("rental")
	fmt.Println(obj.Search("jam") == true)
	fmt.Println(obj.Search("jan") == false)
	fmt.Println(obj.StartsWith("jan") == false)

	runtime.GC()
	runtime.ReadMemStats(&m2)

	fmt.Println()
	fmt.Println("Memory used: ", (m2.Alloc - m1.Alloc), " KB")
	fmt.Println("Total allocated: ", (m2.TotalAlloc - m1.TotalAlloc), " KB")
}
