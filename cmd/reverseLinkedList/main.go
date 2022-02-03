package main

import "fmt"

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
	l1n1, l1n2, l1n3, l1n4, l1n5, l1n6 := ListNode{}, ListNode{}, ListNode{}, ListNode{}, ListNode{}, ListNode{}

	l1n1.Val = 1
	l1n1.Next = &l1n2

	l1n2.Val = 2
	l1n2.Next = &l1n3

	l1n3.Val = 3
	l1n3.Next = &l1n4

	l1n4.Val = 4
	l1n4.Next = &l1n5

	l1n5.Val = 5
	l1n5.Next = &l1n6

	l1n6.Val = 6
	l1n6.Next = nil

	print(&l1n1)
	reverseList(&l1n1)
	print(&l1n1)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack struct {
	list []int
}

func (s *Stack) push(n int) {
	s.list = append(s.list, n)
}

func (s *Stack) peek() int {
	return s.list[len(s.list)-1]
}

func (s *Stack) size() int {
	return len(s.list)
}

func (s *Stack) pop() int {
	removed := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return removed
}

func (s *Stack) print() {
	for i := len(s.list) - 1; i >= 0; i-- {
		fmt.Printf("%v\n", s.list[i])
	}
}

// Given the head of a singly linked list, reverse the list, and return the reversed list.
func reverseList(head *ListNode) *ListNode {
	current := head
	s := Stack{
		list: []int{},
	}
	for current != nil {
		s.push(current.Val)
		current = current.Next
	}
	current = head
	for s.size() != 0 {
		current.Val = s.pop()
		current = current.Next
	}

	return head
}
