package main

import (
	"fmt"

	linkedlist "github.com/mars1385/leetcode/linked_list"
)

func main() {

	h := linkedlist.LinkedList{}

	// for i := 1; i < 6; i++ {
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
	fmt.Println(linkedlist.ReverseList(h.Head))

}
