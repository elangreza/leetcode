package a

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	d := l
	for d != nil {
		fmt.Println(d.Val)
		d = d.Next
	}
}

func rotateRight(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}

	// find deep of list
	deep := 0
	d := dummy.Next
	for d != nil {
		deep++
		d = d.Next
	}

	k = k % deep

	if k == 0 {
		return dummy.Next
	}

	fast := dummy.Next
	slow := dummy.Next
	for range k {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	tail := dummy.Next
	dummy.Next = slow.Next
	fast.Next = tail
	slow.Next = nil

	// dummy.print()

	return dummy.Next
}
