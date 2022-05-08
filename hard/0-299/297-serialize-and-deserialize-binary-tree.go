package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	result := ""

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			result += "#,"
		} else {
			result += fmt.Sprint(node.Val) + ","
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)

	return "[" + result[:len(result)-1] + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	e := strings.Split(data[1:len(data)-1], ",")
	var dfs func(pos int) (*TreeNode, int)

	dfs = func(pos int) (*TreeNode, int) {
		if e[pos] == "#" {
			return nil, pos
		}

		parse, _ := strconv.Atoi(e[pos])
		node := &TreeNode{Val: parse}

		left, next := dfs(pos + 1)
		right, end := dfs(next + 1)
		node.Left = left
		node.Right = right

		return node, end
	}
	root, _ := dfs(0)
	return root
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}}

	ser := Constructor()
	data := ser.serialize(root)
	fmt.Println(data)

	deser := Constructor()
	ans := deser.deserialize(data)
	fmt.Println(ans.Right.Right.Right.Val)
}
