package main_test

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left-right == 0 {
		return head
	}
	h := &ListNode{
		Next: head,
	}
	var cnt int

	before := h
	for head != nil {
		cnt++
		if left <= cnt && cnt <= right {
			start := before
			stack := []*ListNode{}
			for left <= cnt && cnt <= right {
				cnt++
				stack = append(stack, head)
				head = head.Next
			}
			end := head
			for i := len(stack) - 1; i > 0; i-- {
				start.Next = stack[i]
				stack[i].Next = stack[i-1]
				start = stack[i]
			}
			stack[0].Next = end
		}
		before = head
		if head != nil {
			head = head.Next
		}
	}
	return h.Next
}
