package main_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return back(root, 0)
}

func back(node *TreeNode, level int) int {
	if node == nil {
		return level
	}
	l := back(node.Left, level+1)
	r := back(node.Right, level+1)
	return max(l, r)
}
