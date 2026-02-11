package removenthnodefromendoflist

import "fmt"

func (l *ListNode) print(s string) {

	fmt.Println("----", s)
	d := l
	for d != nil {
		fmt.Println(d.Val)
		d = d.Next
	}

	fmt.Println("----", s)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := &ListNode{
		Next: head,
	}
	fast := dummy
	slow := dummy

	for range n {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.print("slow")
	fast.print("fast")

	slow.Next = slow.Next.Next

	return dummy.Next
}
