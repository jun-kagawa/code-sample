package main_test

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	m := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		m[n] = struct{}{}
	}
	dummy := &ListNode{Next: head}
	pre := dummy

	for pre.Next != nil {
		if _, ok := m[pre.Next.Val]; ok {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return dummy.Next
}
