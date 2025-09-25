package main_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	r := []int{}
	back(root, &r)
	return r
}

func back(node *TreeNode, r *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		(*r) = append(*r, node.Val)
		return
	}
	back(node.Left, r)
	(*r) = append(*r, node.Val)
	back(node.Right, r)
}
