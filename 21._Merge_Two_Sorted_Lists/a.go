package mergetwosortedlists

import "fmt"

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	res := &ListNode{}
	current := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = &ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next
		} else {
			current.Next = &ListNode{
				Val: list2.Val,
			}
			list2 = list2.Next
		}

		current = current.Next
	}

	if list2 != nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return res.Next
}
