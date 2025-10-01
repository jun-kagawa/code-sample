package main_test

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	nums := []int{}
	back(root, func(node *TreeNode) {
		nums = append(nums, node.Val)
	})

	slices.Sort(nums)

	var idx int
	back(root, func(node *TreeNode) {
		node.Val = nums[idx]
		idx++
	})
}

func back(root *TreeNode, fn func(node *TreeNode)) {
	if root == nil {
		return
	}
	back(root.Left, fn)
	fn(root)
	back(root.Right, fn)
}
