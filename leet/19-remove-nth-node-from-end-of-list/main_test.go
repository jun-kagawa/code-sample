package main_test

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := head
	limit := n + 1
	arr := make([]*ListNode, 0, limit)
	before := head
	for {
		if len(arr) < limit {
			arr = append(arr, before)
		} else {
			arr = append(arr[1:], before)
		}
		before = before.Next
		if before == nil {
			break
		}
	}
	if len(arr) == 1 {
		return nil
	}
	t := len(arr) - n
	if t - 1 < 0 {
		return arr[t].Next
	}
	if t + 1 > len(arr) - 1 {
		arr[t-1].Next = nil
	} else {
		arr[t-1].Next = arr[t+1]
	}
	return h
}

