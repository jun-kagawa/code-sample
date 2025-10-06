package main_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	m := make(map[int][]int)
	traverse(root, 0, m)
	l := len(m)
	r := [][]int{}
	for i := range l {
		if i%2 == 0 {
			r = append(r, m[i])
		} else {
			r = append(r, reverse(m[i]))
		}
	}
	return r
}

func traverse(node *TreeNode, level int, m map[int][]int) {
	if node == nil {
		return
	}
	m[level] = append(m[level], node.Val)
	traverse(node.Left, level+1, m)
	traverse(node.Right, level+1, m)
}

func reverse(arr []int) []int {
	l := len(arr)
	ar := make([]int, 0, l)
	for i := l - 1; i >= 0; i-- {
		ar = append(ar, arr[i])
	}
	return ar
}
