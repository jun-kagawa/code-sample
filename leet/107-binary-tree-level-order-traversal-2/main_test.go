package main_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	r := [][]int{}
	if root == nil {
		return r
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		level := make([]int, len(queue))
		nextQueue := []*TreeNode{}
		for i, node := range queue {
			level[i] = node.Val
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		r = append(r, []int{})
		copy(r[1:], r)
		r[0] = level

		queue = nextQueue
	}
	return r
}
