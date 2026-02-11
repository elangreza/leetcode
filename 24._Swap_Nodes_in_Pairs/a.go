package swapnodesinpairs

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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  0 1 2 3 4 5
//  3 4 5 rest
//  0 a
//  1 b
//  2 c
//  a -> c -> b -> rest
//  0 2 1 3 4 5
func swapPairs(head *ListNode) *ListNode {
	res := &ListNode{Next: head}
	a := res

	for a.Next != nil && a.Next.Next != nil {
		// init
		b := a.Next
		c := b.Next
		rest := c.Next

		// swap process
		a.Next = c
		c.Next = b
		b.Next = rest

		// next iter set to b
		a = b
	}

	return res.Next
}
