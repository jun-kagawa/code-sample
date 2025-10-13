package main_test

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return back(root, 0) >= 0
}

func back(node *TreeNode, c int) int {
	if node == nil {
		return c
	}
	l := back(node.Left, c + 1)
	r := back(node.Right, c + 1)

	if l == -1 || r == -1 || math.Abs(float64(l - r)) > 1 {
		return -1
	}
	return max(l, r)
}