package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEndV0(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	prt := head
	length := 1
	for prt != nil {
		length++
		prt = prt.Next
	}

	dummy := &ListNode{}
	dummy.Next = head

	length -= n

	curr, prev := dummy, dummy
	for length != 0 {
		if curr != nil {
			prev = curr
			curr = prev.Next
		}
		length--
	}

	next := prev.Next
	prev.Next = next.Next

	return dummy.Next
}

// func print(a string, head *ListNode) {
// 	prt := head
// 	for prt != nil {
// 		fmt.Println(a, prt.Val)
// 		prt = prt.Next
// 	}
// }
