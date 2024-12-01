package linkedlist

import (
	"sort"
)

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	}

	var items []int

	for list1 != nil {

		items = append(items, list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {

		items = append(items, list2.Val)
		list2 = list2.Next
	}

	if len(items) == 0 {
		return nil
	}
	newLinkList := &ListNode{}
	if len(items) == 1 {
		newLinkList.Val = items[0]
		return newLinkList
	}

	sort.Ints(items)

	temp := &ListNode{}

	for index, val := range items {
		if index == 0 {
			newLinkList.Val = val
			newLinkList.Next = temp
		}
		if index != 0 {
			temp.Val = val
		}

		if index != 0 && index < len(items)-1 {
			temp.Next = &ListNode{}
			temp = temp.Next
		}

	}

	return newLinkList
}