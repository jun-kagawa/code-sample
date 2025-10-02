package main_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		if r := isSameTree(p.Left, q.Left); !r {
			return r
		}
		if r := isSameTree(p.Right, q.Right); !r {
			return r
		}
	} else if p == nil && q == nil {
		return true
	} else {
		return false
	}
	return true
}
