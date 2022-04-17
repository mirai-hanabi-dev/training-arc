package main

import "fmt"

type trieNode struct {
	nextChar map[byte]*trieNode
	isEnd    bool
}

type position struct {
	x, y int
}

type queueNode struct {
	x, y    int
	used    map[int]bool
	node    *trieNode
	current []byte
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

func findWords(board [][]byte, words []string) []string {
	n := len(board)
	m := len(board[0])

	var root trieNode
	for _, word := range words {
		root.addWord(word)
	}

	queue := []queueNode{}
	vx := []int{-1, 0, 1, 0}
	vy := []int{0, 1, 0, -1}

	fmt.Println(root.nextChar[111].nextChar[97].nextChar[116].nextChar[104])
	for nextChar := range root.nextChar {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if board[i][j] == nextChar {
					queue = append(queue, queueNode{
						x:       i,
						y:       j,
						used:    map[int]bool{i*m + j: true},
						node:    root.nextChar[nextChar],
						current: []byte{nextChar},
					})
				}
			}
		}
	}

	res := map[string]bool{}

	for len(queue) > 0 {
		fmt.Println(queue)
		nextQueue := []queueNode{}
		for _, q := range queue {
			if q.node.isEnd {
				res[string(q.current)] = true
			}
			for k := 0; k < 4; k++ {
				nx := q.x + vx[k]
				ny := q.y + vy[k]
				if nx == -1 || nx == n {
					continue
				}
				if ny == -1 || ny == m {
					continue
				}
				if _, ok := q.node.nextChar[board[nx][ny]]; ok {
					if _, ok := q.used[nx*m+ny]; !ok {
						node := queueNode{
							x:       nx,
							y:       ny,
							used:    map[int]bool{},
							node:    q.node.nextChar[board[nx][ny]],
							current: append(q.current, board[nx][ny]),
						}
						node.used[nx*m+ny] = true
						for k, v := range q.used {
							node.used[k] = v
						}
						nextQueue = append(nextQueue, node)
					}
				}
			}

		}

		queue = nextQueue
	}

	resArr := []string{}
	for k := range res {
		resArr = append(resArr, k)
	}

	return resArr
}

func main() {
	board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain", "oathi", "oathk", "oathf", "oate", "oathii", "oathfi", "oathfii"}
	fmt.Println(findWords(board, words))
}
