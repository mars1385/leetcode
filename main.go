package main

import (
	"fmt"

	linkedlist "github.com/mars1385/leetcode/linked_list"
)

func main() {

	h := linkedlist.LinkedList{}
	h2 := linkedlist.LinkedList{}

	// for i := 1; i < 3; i++ {
	// 	newNode := &linkedlist.ListNode{
	// 		Val: i,
	// 	}

	// 	if h.Head == nil {
	// 		h.Head = newNode
	// 	} else {

	// 		current := h.Head
	// 		for current.Next != nil {
	// 			current = current.Next
	// 		}
	// 		current.Next = newNode
	// 	}
	// }

	// for i := 1; i < 3; i++ {
	// 	newNode := &linkedlist.ListNode{
	// 		Val: i,
	// 	}

	// 	if h2.Head == nil {
	// 		h2.Head = newNode
	// 	} else {

	// 		current := h2.Head
	// 		for current.Next != nil {
	// 			current = current.Next
	// 		}
	// 		current.Next = newNode
	// 	}
	// }
	d := linkedlist.MergeTwoLists(h.Head, h2.Head)
	for d != nil {

		fmt.Println(d.Val)
		d = d.Next
	}

}
