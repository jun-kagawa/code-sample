package main_test

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var r *ListNode
	if list1.Val < list2.Val {
		r = list1
		list1 = list1.Next
	} else {
		r = list2
		list2 = list2.Next
	}
	point := r
	for list1 != nil || list2 != nil {
		if list1 == nil {
			point.Next = list2
			point = list2
			list2 = list2.Next
		} else if list2 == nil {
			point.Next = list1
			point = list1
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			point.Next = list1
			point = list1
			list1 = list1.Next
		} else {
			point.Next = list2
			point = list2
			list2 = list2.Next
		}
	}
	return r
}
