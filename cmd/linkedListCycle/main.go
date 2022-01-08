package main

import "fmt"

func main() {
	// 3 -> 2 -> 0 -> -4
	//      ^ <- <- <- |

	var third ListNode
	second := ListNode{
		2,
		&third,
	}
	fourth := ListNode{
		-4,
		&second,
	}
	third = ListNode{
		0,
		&fourth,
	}
	first := ListNode{
		3,
		&second,
	}

	// first := ListNode{
	// 	1,
	// 	nil,
	// }
	fmt.Println(hasCycle(&first))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	current := head

	for current.Next != nil {
		cycle := current.Next

		for cycle != nil {
			if cycle == current {
				return true
			}

			cycle = cycle.Next
		}

		current = current.Next
	}

	return false
}

/*

ListNode current // 2

while(current.next != null) {
	ListNode cycle = current.next // 0

	while(cycle != null) {
		if(cycle == current) { // false
			return true
		}

		cycle = cycle.next
	}

	current = current.next
}

return false

*/
