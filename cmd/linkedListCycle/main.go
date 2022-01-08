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
	current := head
	visited := []*ListNode{}

	for current != nil {

		for i := 0; i < len(visited); i++ {
			if visited[i] == current {
				return true
			}
		}

		visited = append(visited, current)

		current = current.Next
	}

	return false
}

/*

ListNode current // 3 // 2 // 0 // -4 // 2
ListNode[] visited // [] // [3] \\ [3, 2] // [3, 2, 0] // [3, 2, 0, -4]

while(current != null) {

	for i = 0; i < visited.length; i++ { //
		if visited[i] == current {
			return true
		}
	}
	visited.append(current)

	current = curren.next
}

return false










*/
