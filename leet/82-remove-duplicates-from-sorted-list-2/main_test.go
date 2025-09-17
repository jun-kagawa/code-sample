package main_test

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	h := &ListNode{
		Val:  0,
		Next: head,
	}
	prev := h
	curr := head

	for curr != nil {
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr = curr.Next
		}

		if prev.Next == curr {
			prev = prev.Next
		} else {
			prev.Next = curr.Next
		}
		curr = curr.Next
	}
	return h.Next
}
