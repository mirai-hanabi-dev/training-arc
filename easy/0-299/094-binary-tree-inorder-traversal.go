package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func travel(node *TreeNode, order *[]int) {
	if node == nil {
		return
	}
	travel(node.Left, order)
	*order = append(*order, node.Val)
	travel(node.Right, order)
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	travel(root, &result)
	return result
}
