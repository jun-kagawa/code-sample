package main_test

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	h := head.Next
	if h == nil {
		return head
	}
	var t *ListNode = nil
	for {
		b := head
		if b == nil {
			break
		}
		n := head.Next
		if n == nil {
			break
		}
		b.Next = n.Next
		n.Next = b
		head = b.Next
		if t != nil {
			t.Next = n
		}
		t = b
	}

	return h
}
