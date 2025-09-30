package main_test

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	maxx := math.MaxInt
	minn := math.MinInt
	return recur(root, minn, maxx)
}

func recur(cur *TreeNode, minn, maxx int) bool {
	if cur == nil {
		return true
	}
	v := cur.Val
	if cur.Left != nil {
		l := cur.Left.Val
		if minn < l && l < v {
			if r := recur(cur.Left, minn, v); !r {
				return false
			}
		} else {
			return false
		}
	}

	if cur.Right != nil {
		r := cur.Right.Val
		if v < r && r < maxx {
			if r := recur(cur.Right, v, maxx); !r {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
