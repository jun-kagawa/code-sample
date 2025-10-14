package main_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	return back(root, 0)
}

func back(node *TreeNode, c int) int {
	if node == nil {
		return -1
	}
	if node.Left == nil && node.Right == nil {
		return c
	}
	l := back(node.Left, c+1)
	r := back(node.Right, c+1)
	if l == -1 {
		return r
	} else if r == -1 {
		return l
	} else {
		return min(l, r)
	}
}
