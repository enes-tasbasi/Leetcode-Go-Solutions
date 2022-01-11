package main

import "fmt"

func main() {
	first, second, third, fourth, fifth, sixth, seventh := ListNode{}, ListNode{}, ListNode{}, ListNode{}, ListNode{}, ListNode{}, ListNode{}

	first.Val = 1
	first.Next = &second

	second.Val = 2
	second.Next = &third

	third.Val = 3
	third.Next = &fourth

	fourth.Val = 3
	fourth.Next = &fifth

	fifth.Val = 4
	fifth.Next = &sixth

	sixth.Val = 4
	sixth.Next = &seventh

	seventh.Val = 5
	seventh.Next = nil

	printList(deleteDuplicates(&first))
}

func printList(n *ListNode) {
	if n == nil {
		return
	}
	fmt.Printf("%d -> ", n.Val)
	printList(n.Next)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// [1,2,3,3,4,4,5]
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	current := head
	ahead := head.Next // 2 // 3 -> 3 // 3 -> 4 // 4 -> 4 // 4 -> 5

	for ahead != nil {
		if current.Next.Val == ahead.Val {
			ahead = ahead.Next
			if ahead != nil {
				ahead = ahead.Next
			}
			current.Next = ahead
		} else {
			current = ahead
			ahead = ahead.Next
		}
	}

	return head
}

/*

// 1 -> 2 -> 3 -> 3

	     ahead
			|
1  			2	-> 3 -> 3
|			^
- - - - - - |


			 			ahead/current

	2	-> 3   -> 3  	-> nil



if(head == null) {
	return null
}

current = head
ahead = head.next

while(ahead != null) {
	if(current.val == ahead.val) {
		ahead = ahead.next
		current = ahead

		if(ahead != null) {
			ahead = ahead.next
		}
	} else {
		ahead = ahead.next
		current = current.next
	}
}

return head





*/
