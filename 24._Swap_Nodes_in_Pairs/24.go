package swapnodesinpairs

import "fmt"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 1 2 3 4 5 6
// c d

// c d
// curr.next = d
// c.next = d.Next
// d.next = c

// 2 1 4 3 6 5

func swapPairsV0(head *ListNode) *ListNode {
	if head != nil && head.Next == nil {
		return head
	}

	res := &ListNode{}
	curr := res
	for head != nil && head.Next != nil {
		// for swapping position
		a := head.Next
		head.Next = head.Next.Next
		a.Next = head

		// for assigning into res new head
		curr.Next = a
		curr = curr.Next

		// for assigning into res new head.next
		curr.Next = a.Next
		curr = curr.Next

		// skipping into next
		// why only skip once
		// since swap in the logic
		// we got directly even (n) position of listNode
		head = head.Next
	}

	return res.Next
}

func print(a string, head *ListNode) {
	prt := head
	for prt != nil {
		fmt.Println(a, prt.Val)
		prt = prt.Next
	}
}
