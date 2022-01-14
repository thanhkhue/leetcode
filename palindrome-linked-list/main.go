// https://leetcode.com/problems/palindrome-linked-list/
package main

import (
	"github/leetcode/utils"
	"log"
)

func isPalindrome(head *utils.ListNode) bool {
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	list := utils.ArrayToLinkedList(arr)
	noExtraSpaceSol(list)
}

func noExtraSpaceSol(head *utils.ListNode) bool {
	head1 := head

	// find middle
	slow, fast := head1, head1
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow
	log.Printf("Middle item %d", mid.Val)

	// reverse haft linkedlist from mid to end
	revList := utils.ReverseLinkedList(slow)

	// compare
	for revList.Next != nil {
		log.Printf("%d vs %d", revList.Val, head.Val)
		if revList.Val != head.Val {
			return false
		}
		revList = revList.Next
		head = head.Next
	}

	return true
}

func usingArray(head *utils.ListNode) bool {
	arr := make([]*utils.ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	for i := 0; i < len(arr)/2; i++ {
		if arr[i].Val != arr[len(arr)-i-1].Val {
			return false
		}
	}
	return true
}
