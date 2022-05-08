package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursive(preorder []int, inorder []int, rootIndex *int, mappingInorder map[int]int, startInorder int, endInorder int) *TreeNode {
	if startInorder > endInorder {
		return nil
	}
	val := preorder[*rootIndex]
	node := TreeNode{
		Val: val,
	}

	midInorder := mappingInorder[val]

	*rootIndex++
	node.Left = recursive(preorder, inorder, rootIndex, mappingInorder, startInorder, midInorder-1)
	node.Right = recursive(preorder, inorder, rootIndex, mappingInorder, midInorder+1, endInorder)

	return &node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	mappingInorder := map[int]int{}
	for idx, value := range inorder {
		mappingInorder[value] = idx
	}
	rootIndex := 0
	return recursive(preorder, inorder, &rootIndex, mappingInorder, 0, len(inorder)-1)
}
