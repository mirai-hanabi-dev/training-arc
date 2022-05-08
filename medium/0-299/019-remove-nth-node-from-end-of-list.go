package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	second := head
	tmp := n
	for tmp > 0 {
		second = second.Next
		tmp--
	}
	var prev *ListNode
	first := head
	for second != nil {
		if prev == nil {
			prev = first
		} else {
			prev = prev.Next
		}
		first = first.Next
		second = second.Next
	}

	if prev == nil {
		return head.Next
	}
	prev.Next = first.Next
	return head
}
