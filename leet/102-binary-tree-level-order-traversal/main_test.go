package main_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	r := make(map[int][]int)
	traverse(root, 0, r)

	l := len(r)
	if l == 0 {
		return [][]int{}
	}
	lst := [][]int{}
	for i := range l {
		lst = append(lst, r[i])
	}
	return lst
}

func traverse(node *TreeNode, level int, m map[int][]int) {
	if node == nil {
		return
	}
	m[level] = append(m[level], node.Val)

	traverse(node.Left, level+1, m)
	traverse(node.Right, level+1, m)
}
