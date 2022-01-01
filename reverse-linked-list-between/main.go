package main

import (
	"fmt"
	"log"
)

// https://leetcode.com/problems/reverse-linked-list-ii/
// follow-up: Could you do it in one pass?

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	linkedList := arrayToLinkedList(arr)
	r := reverseBetweenInOnePass(linkedList, 2, 8)
	fmt.Printf("Reverse result %+v", linkedListToArray(r))
	// reverseBetween(linkedList, 2, 4)
}

func reverseBetweenInOnePass(head *ListNode, left, right int) *ListNode {
	if left == right {
		return head
	}

	dummy := head
	for i := 0; i < left; i++ {
		head = head.Next
	}

	prev := (*ListNode)(nil)
	cur := head.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next // store next pointer
		cur.Next = prev  // reverse happens
		prev = cur
		cur = next
	}

	head.Next.Next = cur
	head.Next = prev
	return dummy
}

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	next.Next = head
	return next
}

func linkedListToArray(head *ListNode) []int {
	arr := make([]int, 0)
	for head.Next != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return append(arr, head.Val)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	arr := linkedListToArray(head)
	arr = changePositions(arr, left, right)
	log.Printf("Reverse array %+v", arr)
	return arrayToLinkedList(arr)
}

func changePositions(arr []int, leftPos int, rightPos int) []int {
	if rightPos > len(arr) {
		return arr
	}

	for leftPos < rightPos {
		arr[leftPos-1], arr[rightPos-1] = arr[rightPos-1], arr[leftPos-1]
		leftPos++
		rightPos--
	}

	return arr
}

func foundBoth(head *ListNode, left int, right int) bool {
	p := head
	foundLeft, foundRight := false, false
	for p.Next != nil {
		if p.Val == left {
			foundLeft = true
		}
		if p.Val == right {
			foundRight = true
		}
		p = p.Next
	}
	if p.Val == left {
		foundLeft = true
	}
	if p.Val == right {
		foundRight = true
	}

	log.Printf("Found hay la ko: left %t, right %t", foundLeft, foundRight)

	return foundLeft && foundRight
}

func newNode(n *ListNode, val int) *ListNode {
	if n == nil {
		return &ListNode{
			Val: val,
		}
	}

	if n.Next == nil {
		n.Next = &ListNode{
			Val: val,
		}
		return n
	}

	p := n
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &ListNode{Val: val}
	return n
}

func arrayToLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	n := &ListNode{
		Val: arr[0],
	}
	for i := 1; i < len(arr); i++ {
		n = newNode(n, arr[i])
	}
	return n
}
