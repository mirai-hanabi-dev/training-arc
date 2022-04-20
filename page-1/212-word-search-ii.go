package main

type trieNode struct {
	nextChar map[byte]*trieNode
	isEnd    bool
}

type position struct {
	x, y int
}

func (t *trieNode) addWord(word string) {
	cur := t
	for _, c := range word {
		if cur.nextChar == nil {
			cur.nextChar = make(map[byte]*trieNode)
		}
		if _, ok := cur.nextChar[byte(c)]; !ok {
			cur.nextChar[byte(c)] = &trieNode{}
		}
		cur = cur.nextChar[byte(c)]
	}
	cur.isEnd = true
}

func dfs(i, j int, board *[][]byte, node *trieNode, curRes string, finalRes *[]string) {
	if node.isEnd {
		*finalRes = append(*finalRes, curRes)
		node.isEnd = false
	}
	if i < 0 || i == len(*board) {
		return
	}
	if j < 0 || j == len((*board)[0]) {
		return
	}
	if _, ok := node.nextChar[(*board)[i][j]]; !ok {
		return
	}
	char := (*board)[i][j]
	(*board)[i][j] = '#'

	vx := []int{-1, 0, 1, 0}
	vy := []int{0, 1, 0, -1}
	for k := 0; k < 4; k++ {
		dfs(i+vx[k], j+vy[k], board, node.nextChar[char], curRes+string(char), finalRes)
	}

	(*board)[i][j] = char
}

func findWords(board [][]byte, words []string) []string {
	var root trieNode
	for _, word := range words {
		root.addWord(word)
	}

	resArr := []string{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, &board, &root, "", &resArr)
		}
	}

	return resArr
}
