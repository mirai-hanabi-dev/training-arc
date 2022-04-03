package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, result *int) int {
	leftPath := math.MinInt32
	if root.Left != nil {
		leftPath = dfs(root.Left, result)
	}

	rightPath := math.MinInt32
	if root.Right != nil {
		rightPath = dfs(root.Right, result)
	}

	best := root.Val
	best = max(best, root.Val+leftPath)
	best = max(best, root.Val+rightPath)

	*result = max(*result, best)
	*result = max(*result, root.Val+leftPath+rightPath)

	return best
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathSum(root *TreeNode) int {
	result := math.MinInt32

	dfs(root, &result)

	return result
}
