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
	cycle := detectCycle(&first)
	if cycle != nil {
		fmt.Println(cycle)
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func interception(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return slow
		}
	}

	return nil
}

func detectCycle(head *ListNode) *ListNode {
	start := head // 1
	intercept := interception(head)

	if intercept == nil {
		return nil
	}

	for start != intercept {
		start = start.Next
		intercept = intercept.Next
	}

	return start
}

/*



fast = head
slow = head

while fast != null && fast.Next != null {
	fast = fast.Next.Next
	slow = slow.Next

	if(fast == slow) {
		return
	}
}

return null












*/
