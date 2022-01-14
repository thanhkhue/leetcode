// https://leetcode.com/problems/rotate-list/

package main

import (
	"github/leetcode/utils"
	"log"
)

/*
1 -> 2 -> 3 -> 4 -> 5
if k = 2 {
	4 -> 5 -> 1 -> 2 -> 3
}
*/

func rotateRight(head *utils.ListNode, k int) *utils.ListNode {
	if k == 0 {
		return head
	}
	if head == nil || head.Next == nil {
		return head
	}

	head1 := head
	l := 0
	for head1 != nil {
		l++
		head1 = head1.Next
	}

	k = k % l
	if k == 0 {
		return head
	}

	fast := head
	slow := head
	log.Printf("slow %d", slow.Val)
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	log.Printf("fast %d", fast.Val)

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	log.Printf("slow %d", slow.Val)
	log.Printf("fast %d", fast.Val)
	nextSlow := slow.Next
	slow.Next = nil
	fast.Next = head

	return nextSlow
}

func rotateRightV1(head *utils.ListNode, k int) *utils.ListNode {
	head1 := head
	l := 0
	for head1 != nil {
		l++
		head1 = head1.Next
	}
	log.Printf("l %d, k=%d", l, k%l)

	k = k % l
	if k == 0 {
		return head
	}

	if l == 2 {
		head2 := head
		if k == 2 {
			return head
		}
		if k == 1 {

			head.Val, head.Next.Val = head.Next.Val, head.Val
		}
		log.Printf("head %d -> %d", head2.Val, head2.Next.Val)
		return head2
	}

	slow, fast := head, head.Next
	for k > 0 && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		k--
	}

	log.Printf("slow %d", slow.Val)
	log.Printf("fast %d", fast.Val)

	slow.Next = nil
	head2 := fast
	for fast.Next != nil {
		fast = fast.Next
	}
	fast.Next = head

	log.Printf("fast %+v", fast)

	log.Printf("head2 %d ", head2.Val)
	for head2 != nil {
		log.Printf("head2.Val %d", head2.Val)
		head2 = head2.Next
	}
	return head2
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	list1 := utils.ArrayToLinkedList(arr)
	log.Println("reach here")
	rotateRight(list1, 2)

	// -> [2, 1]
}
