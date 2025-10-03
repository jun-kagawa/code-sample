package main_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return symmetric(root.Left, root.Right)
}

func symmetric(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l != nil && r != nil && l.Val == r.Val {
		if r := symmetric(l.Left, r.Right); !r {
			return false
		}
		if r := symmetric(l.Right, r.Left); !r {
			return false
		}
		return true
	}
	return false
}
