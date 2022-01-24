package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func recursePrint(node *ListNode) {
	if node == nil {
		return
	}

	if node.Next != nil {
		fmt.Printf("%v -> ", node.Val)
	} else {
		fmt.Printf("%v", node.Val)
	}

	recursePrint(node.Next)
}

func print(node *ListNode) {
	recursePrint(node)
	fmt.Println()
}

func main() {
	l1n1, l1n2, l1n3, l2n1, l2n2, l2n3 := ListNode{}, ListNode{}, ListNode{}, ListNode{}, ListNode{}, ListNode{}

	l1n1.Val = 2
	l1n1.Next = &l1n2

	l1n2.Val = 4
	l1n2.Next = &l1n3

	l1n3.Val = 3
	l1n3.Next = nil

	l2n1.Val = 5
	l2n1.Next = &l2n2

	l2n2.Val = 6
	l2n2.Next = &l2n3

	l2n3.Val = 4
	l2n3.Next = nil

	print(&l1n1)
	print(&l2n1)
	print(addTwoNumbers(&l1n1, &l2n1))

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	current1 := l1
	current2 := l2

	sumHead := &ListNode{}
	sumCurrent := sumHead
	carry := 0

	for current1 != nil && current2 != nil {
		newSum := current1.Val + current2.Val + carry
		carry = 0

		if newSum > 9 {
			carry = 1
			sumCurrent.Val = newSum % 10
		} else {
			sumCurrent.Val = newSum
		}

		if current1.Next != nil {
			sumCurrent.Next = &ListNode{}
			sumCurrent = sumCurrent.Next
		}

		current1 = current1.Next
		current2 = current2.Next
	}

	if carry > 0 {
		sumCurrent.Next = &ListNode{Val: carry}
	}

	return sumHead
}

/*

2 -> 4 -> 3  // 342
5 -> 6 -> 4 // 465

l1, l2

current1 = l1
current2 = l2
sumHead = ListNode
sumCurrent = ListNode
carry = 0

while current1 != null && current2 != null {

	newSum = current1.val + current2.val + carry
	carry = 0

	if newSum > 9 {
		carry = 1

		sumCurrent.val = newSum % 10
	} else {
		sumCurrent.val = newSum
	}

	sumCurrent.next = ListNode{}
	sumCurrent = sumCurrent.next

}

if carry > 0 {
	sumCurrent.next = ListNode{val: carry}
}

return sumHead









*/
