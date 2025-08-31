package main_test

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}
	c := head
	var l int = 1
	for c.Next != nil {
		c = c.Next
		l++
	}
	cc := k % l
	var before *ListNode
	current := head
	for range cc {
		for current.Next != nil {
			before = current
			current = current.Next
		}
		if before != nil {
			before.Next = nil
		}
		current.Next = head
		head = current
	}
	return current
}
