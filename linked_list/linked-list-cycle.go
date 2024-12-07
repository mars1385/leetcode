package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func HasCycle(head *ListNode) bool {

	posLoc := map[*ListNode]struct{}{}

	for head != nil {

		_, hash := posLoc[head]
		if hash {
			return true
		}
		posLoc[head] = struct{}{}

		head = head.Next
	}

	return false
}
