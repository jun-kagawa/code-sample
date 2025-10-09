package main_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	m := make(map[int]int, len(inorder))
	for i, n := range inorder {
		m[n] = i
	}

	i := len(postorder) - 1
	var helper func(j, k int) *TreeNode
	helper = func(j, k int) *TreeNode {
		if j > k {
			return nil
		}
		val := postorder[i]
		i--
		idx := m[val]
		node := &TreeNode{Val: val}
		node.Right = helper(idx+1, k)
		node.Left = helper(j, idx-1)
		return node
	}
	return helper(0, len(postorder)-1)
}
