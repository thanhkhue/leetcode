package main

import (
	"fmt"
	"log"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
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

	p.Next = &ListNode{
		Val: val,
	}
	return n
}

func reverseList(head *ListNode) *ListNode {
	return linkedListFromArray(linkedListToArray(head))
}

func linkedListFromArray(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	n := &ListNode{
		Val: arr[len(arr)-1],
	}
	for i := len(arr) - 2; i >= 0; i-- {
		n = newNode(n, arr[i])
	}

	return n
}

func linkedListToArray(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	arr := make([]int, 0)
	for head.Next != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	if head == nil {
		return arr
	}

	return append(arr, head.Val)
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

func printLinkedList(n *ListNode) {
	p := n
	for p.Next != nil {
		fmt.Printf("%d -> ", p.Val)
		p = p.Next
	}

	if p != nil {
		fmt.Printf("%d", p.Val)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	n := arrayToLinkedList(arr)
	printLinkedList(n)
	log.Println("====================")
	printLinkedList(reverseList(n))
}
