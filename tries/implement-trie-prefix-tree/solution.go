package implementtrieprefixtree

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

type Trie struct {
	root *node
}

func Constructor() Trie {
	return Trie{root: newNode()}
}

func (t *Trie) Insert(word string) {
	curr := t.root

	for _, ch := range word {
		key := string(ch)
		if _, ok := curr.children[key]; !ok {
			curr.children[key] = newNode()
		}

		curr = curr.children[key]
	}

	curr.eow = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root

	for _, ch := range word {
		key := string(ch)
		if _, ok := curr.children[key]; !ok {
			return false
		}

		curr = curr.children[key]
	}

	return curr.eow
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root

	for _, ch := range prefix {
		key := string(ch)
		if _, ok := curr.children[key]; !ok {
			return false
		}

		curr = curr.children[key]
	}

	return true
}
