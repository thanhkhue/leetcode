/*
https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{1}
	linkedList := arrayToLinkedList(arr)
	fmt.Printf("Results %+v", linkedListToArray(deleteMiddle(linkedList)))
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return nil
	}
	middle := findTheMiddle(head)
	nextMid := middle.Next
	dummy := head
	for head.Next != nil {
		if head.Next == middle {
			head.Next = nextMid
		} else {
			head = head.Next
		}
	}
	return dummy
}

func findTheMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
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
