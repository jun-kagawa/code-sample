package main_test

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int, len(inorder))
	for i, n := range inorder {
		m[n] = i
	}
	var i int

	var helper func(j, k int) *TreeNode
	helper = func(j, k int) *TreeNode {
		if j > k {
			return nil
		}
		val := preorder[i]
		i++
		node := &TreeNode{Val: val}
		idx := m[val]
		node.Left = helper(j, idx-1)
		node.Right = helper(idx+1, k)
		return node
	}
	return helper(0, len(inorder)-1)
}
