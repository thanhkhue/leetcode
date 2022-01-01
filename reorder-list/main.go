/*
https://leetcode.com/problems/reorder-list/
*/

package main

import (
	"fmt"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Results %+v", linkedListToArray(reorderInplace(arrayToLinkedList(arr))))
}

func reorderInplace(head *ListNode) *ListNode {
	mid := findMiddleItem(head)
	midNext := mid.Next
	mid.Next = nil
	log.Printf("Mid %d", mid.Val)
	firstHalf, secondHalf := head, reverseSecondHalf(midNext)
	rearrangeList := new(ListNode)
	// rearrangeList.Next = nil
	fmt.Printf("first half %+v", linkedListToArray(firstHalf))
	dummy := rearrangeList
	for firstHalf != nil && secondHalf != nil {
		if firstHalf != nil {
			log.Printf("First %d", firstHalf.Val)
			rearrangeList.Next = firstHalf
			rearrangeList = rearrangeList.Next
			firstHalf = firstHalf.Next
		}
		if secondHalf != nil {
			log.Printf("Second %d", secondHalf.Val)
			rearrangeList.Next = secondHalf
			rearrangeList = rearrangeList.Next
			secondHalf = secondHalf.Next
		}
	}
	if firstHalf != nil {
		rearrangeList.Next = firstHalf
	}
	if secondHalf != nil {
		rearrangeList.Next = secondHalf
	}
	dummy = dummy.Next

	return dummy
}

func reverseSecondHalf(mid *ListNode) *ListNode {
	if mid == nil {
		return mid
	}

	prev, cur, next := (*ListNode)(nil), mid, (*ListNode)(nil)
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func findMiddleItem(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	log.Printf("slow %d", slow.Val)
	return slow
}

func usingArray(head *ListNode) *ListNode {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	log.Printf("First item %d", arr[0])

	left, right := 1, len(arr)-1
	rearrange := &ListNode{
		Val:  arr[0],
		Next: new(ListNode),
	}
	first := rearrange
	for i := 0; i < len(arr)-1; i++ {
		if i%2 != 0 {
			rearrange.Next = &ListNode{
				Val: arr[left],
				// Next: new(ListNode),
			}
			log.Printf("odd val %d", rearrange.Next.Val)
			left++
		} else {
			rearrange.Next = &ListNode{
				Val: arr[right],
				// Next: new(ListNode),
			}
			log.Printf("Even Val %d", rearrange.Next.Val)
			right--
		}
		rearrange = rearrange.Next
	}
	rearrange = nil

	return first
}

func reverse(head *ListNode) *ListNode {
	prev, cur := (*ListNode)(nil), head
	for head != nil {
		next := head.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
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
