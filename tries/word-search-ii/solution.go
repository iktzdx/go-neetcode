package wordsearchii

type trieNode struct {
	children [26]*trieNode
	word     string
}

func (n *trieNode) insert(word string) {
	for _, ch := range word {
		idx := ch - 'a'
		if n.children[idx] == nil {
			n.children[idx] = new(trieNode)
		}

		n = n.children[idx]
	}

	n.word = word
}

func FindWords(board [][]byte, words []string) []string {
	result := make([]string, 0, len(words))

	trie := new(trieNode)
	for _, word := range words {
		trie.insert(word)
	}

	for i, row := range board {
		for j, ch := range row {
			idx := ch - 'a'
			if tn := trie.children[idx]; tn != nil {
				dfs(board, i, j, trie, &result)
			}
		}
	}

	return result
}

func dfs(board [][]byte, i, j int, node *trieNode, result *[]string) {
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	letter := board[i][j]
	node = node.children[letter-'a']

	if node == nil {
		return
	} else if node.word != "" {
		*result = append(*result, node.word)
		node.word = ""
	}

	board[i][j] = '.'

	for _, dir := range directions {
		x, y := i+dir[0], j+dir[1]
		if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && board[x][y] != '.' {
			dfs(board, x, y, node, result)
		}
	}

	board[i][j] = letter
}
