package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var hasTargetNode func(node *TreeNode) bool
	var res *TreeNode

	hasTargetNode = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		left := hasTargetNode(node.Left)
		right := hasTargetNode(node.Right)
		mid := node == p || node == q

		// If 2 out of 3 cases above is correct,
		// the current node is LCA.
		check := []bool{left, mid, right}
		count := 0
		for _, b := range check {
			if b {
				count += 1
			}
		}
		if count == 2 && res == nil {
			res = node
		}

		return mid || left || right
	}

	hasTargetNode(root)

	return res
}
