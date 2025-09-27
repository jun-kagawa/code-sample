package main_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	var trees []*TreeNode
	for i := start; i <= end; i++ {
		left := generate(start, i - 1)
		right := generate(i+1, end)

		for _, l := range left {
			for _, r := range right {
				tree := &TreeNode{Val:i, Left: l, Right: r}
				trees = append(trees, tree)
			}
		}
	}
	return trees
}
