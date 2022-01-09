package main

import "fmt"

func main() {
	// 3 -> 2 -> 0 -> -4
	//      ^ <- <- <- |

	// var third ListNode
	// second := ListNode{
	// 	2,
	// 	&third,
	// }
	// fourth := ListNode{
	// 	-4,
	// 	&second,
	// }
	// third = ListNode{
	// 	0,
	// 	&fourth,
	// }
	// first := ListNode{
	// 	3,
	// 	&second,
	// }

	first := ListNode{
		1,
		nil,
	}
	fmt.Println(hasCycle(&first))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}

	return false
}

/*


 */
