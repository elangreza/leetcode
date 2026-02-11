package mergetwosortedlists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoListsV0(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	curr := res

	// we want to compare between list one and list two
	// if list1 val is lower than list2 val
	// assign lowest val into res Node
	// and iterate until either list1 or list2 is empty

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			curr.Next = list2
			curr = curr.Next
			list2 = list2.Next
		} else {
			curr.Next = list1
			curr = curr.Next
			list1 = list1.Next
		}
	}

	// last step is assigning 1 more node since list one or list two will carry extra 1 value
	// and assign into curr.Next
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
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
