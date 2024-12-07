package main

import (
	"fmt"

	linkedlist "github.com/mars1385/leetcode/linked_list"
)

func main() {

	h := linkedlist.LinkedList{}
	v := []int{3, 2, 0, -4}

	var aa *linkedlist.ListNode
	for i := 0; i <= 3; i++ {
		newNode := &linkedlist.ListNode{
			Val: v[i],
		}

		if i == 1 {
			aa = newNode
		}
		if h.Head == nil {
			h.Head = newNode
		} else {

			current := h.Head
			for current.Next != nil {
				current = current.Next
			}
			current.Next = newNode
		}

		if i == 3 {

			current := h.Head
			for current.Next != nil {
				current = current.Next
			}
			current.Next = aa
		}
	}

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

	d := linkedlist.HasCycle(h.Head)

	fmt.Println(d)

}
