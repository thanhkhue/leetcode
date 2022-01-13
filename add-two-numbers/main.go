// https://leetcode.com/problems/add-two-numbers/

package main

import (
	"../utils"
)

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {

	newList := &utils.ListNode{
		Next: new(utils.ListNode),
	}
	head := newList
	c := 0
	for c > 0 || l1 != nil || l2 != nil {

		if l1 != nil {
			c += l1.Val
		}
		if l2 != nil {
			c += l2.Val
		}
		newList.Next = &utils.ListNode{
			Val: c % 10,
		}
		newList = newList.Next
		c = c / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	head = head.Next
	return head
}

func main() {
	arr1 := []int{2, 4, 3}
	arr2 := []int{5, 6, 4}

	// arr1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	// arr2 := []int{5, 6, 4}
	l1 := utils.ArrayToLinkedList(arr1)
	l2 := utils.ArrayToLinkedList(arr2)
	addTwoNumbers(l1, l2)
}

func reverseLinkedList(l *utils.ListNode) *utils.ListNode {
	prev := (*utils.ListNode)(nil)
	cur := l
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
