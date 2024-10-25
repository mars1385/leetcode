package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func ReverseList(head *ListNode) *ListNode {

	if head == nil {

		return head
	}
	newLinkList := &ListNode{
		Val: head.Val,
	}
	head = head.Next
	for head != nil {

		temp := ListNode{
			Next: newLinkList,
			Val:  head.Val,
		}

		newLinkList = &temp

		head = head.Next
	}

	return newLinkList
}
