package designaddandsearchwordsdatastructure

type node struct {
	children map[string]*node
	eow      bool
}

func newNode() *node {
	return &node{
		children: make(map[string]*node),
		eow:      false,
	}
}

type WordDictionary struct {
	root *node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: newNode(),
	}
}

func (dict *WordDictionary) AddWord(word string) {
	curr := dict.root

	for _, ch := range word {
		key := string(ch)
		if _, ok := curr.children[key]; !ok {
			curr.children[key] = newNode()
		}

		curr = curr.children[key]
	}

	curr.eow = true
}

func (dict *WordDictionary) Search(word string) bool {
	return dfs(dict.root, word, 0)
}

func dfs(root *node, word string, start int) bool {
	curr := root

	for i := start; i < len(word); i++ {
		ch := word[i]

		if ch == '.' {
			for _, child := range curr.children {
				if dfs(child, word, i+1) {
					return true
				}
			}

			return false
		}

		key := string(ch)
		if _, ok := curr.children[key]; !ok {
			return false
		}

		curr = curr.children[key]
	}

	return curr.eow
}
