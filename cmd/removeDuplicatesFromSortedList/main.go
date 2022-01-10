package main

import "fmt"

func main() {
	first, second, third, fourth, fifth := ListNode{}, ListNode{}, ListNode{}, ListNode{}, ListNode{}

	first.Val = 1
	first.Next = &second

	second.Val = 1
	second.Next = &third

	third.Val = 2
	third.Next = &fourth

	fourth.Val = 3
	fourth.Next = &fifth

	fifth.Val = 3
	fifth.Next = nil

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

// Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
// Return the linked list sorted as well.
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	ahead := head.Next

	for ahead != nil {
		if current.Val == ahead.Val {
			ahead = ahead.Next
			current.Next = ahead
		} else {
			ahead = ahead.Next
			current = current.Next
		}
	}

	return head
}

/*
1 -> 1 -> 2 -> 3 -> 3
1 -> 2 -> 3 -> 3

if head.next == null {
	return head
}

current = head | 1 | 2 | 3
ahead = head.next | 1 | 3 | 3

while next != null {
	if current.Val == ahead.Val { | true | false | true
		current.Next = ahead.Next  | current.Next = nil
		ahead = ahead.Next | ahead = nil
	}
	current = current.Next | current = nil
	ahead = ahead.Next
}

return head



1 -> 1 -> 2 -> 3 -> 3
1 -> 2 -> 3 -> 3
1 -> 2 -> 3 -> nil

current = head // 1 // 2 // 3
ahead = head.next // 1 // 2 // 3 // 3 // nil

while next != null {
	if current.Val == ahead.Val { // true // false // false
		ahead = ahead.Next // nil
		current.Next = ahead //
	} else {
		ahead = ahead.Next
		current = current.Next
	}
}

return head




*/
