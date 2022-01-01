package main

import (
	"fmt"

	"../utils"
)

func swapNodes(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	start := head
	for i := 1; i < k; i++ {
		head = head.Next
	}
	left := head
	slow, fast := start, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	right := slow

	left.Val, right.Val = right.Val, left.Val

	return start
}

func main() {
	head := utils.ArrayToLinkedList([]int{80, 46, 66, 35, 64})
	k := 1
	swapped := swapNodes(head, k)
	fmt.Printf("%+v", utils.LinkedListToArray(swapped))
}

func usingArray(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil {
		return head
	}

	arr := utils.LinkedListToArray(head)
	arr[k-1], arr[len(arr)-k] = arr[len(arr)-k], arr[k-1]
	return utils.ArrayToLinkedList(arr)
}

/*
func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	start := head
	prevLeft := head
	for i := 1; i < k; i++ {
		prevLeft = head
		head = head.Next
	}
	left := head
	slow, fast := start, head
	prevRight := slow
	for fast.Next != nil {
		prevRight = slow
		slow = slow.Next
		fast = fast.Next
	}
	right := slow

	if prevLeft == prevRight {
		left.Val, right.Val = right.Val, left.Val
		return start
	}

	prevLeft.Next, prevRight.Next = right, left
	right.Next, left.Next = left.Next, right.Next

	return start
}
*/
