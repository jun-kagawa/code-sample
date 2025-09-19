package main_test

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition2(head *ListNode, x int) *ListNode {
	h := &ListNode{
		Next: head,
	}
	prev := h
	for head != nil {
		if head.Val < x {
			prev = head
			head = head.Next
			continue
		}
		tprev := head
		target := head.Next
		for target != nil && target.Val >= x {
			tprev = target
			target = target.Next
		}
		if target == nil {
			return h.Next
		}
		prev.Next = target
		prev = target
		tprev.Next = target.Next
		target.Next = head
		head = target.Next
	}
	return h.Next
}
