// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/

package main

import (
	"github/leetcode/utils"
	"log"
	"math"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *utils.ListNode) int {
	if head == nil {
		return 0
	}

	if head != nil && head.Next != nil && head.Next.Next == nil {
		return head.Val + head.Next.Val
	}

	slow, fast := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	revList := utils.ReverseLinkedList(slow)
	max := 0
	for revList.Next != nil {
		log.Printf("head.Val + revList.Val = %d", head.Val+revList.Val)
		if head.Val+revList.Val > max {
			max = head.Val + revList.Val
		}
		head = head.Next
		revList = revList.Next
	}
	log.Printf("max %d", max)
	return max
}

func pairSumUsingArray(head *utils.ListNode) int {
	if head == nil {
		return 0
	}

	if head != nil && head.Next != nil && head.Next.Next == nil {
		return head.Val + head.Next.Val
	}

	arr := make([]*utils.ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	max := math.MinInt
	for i := 0; i < len(arr)/2; i++ {
		log.Printf("sum %d", arr[i].Val+arr[len(arr)-i-1].Val)
		if arr[i].Val+arr[len(arr)-i-1].Val > max {
			max = arr[i].Val + arr[len(arr)-i-1].Val
		}
	}

	log.Printf("max %d", max)
	return max
}

func main() {
	arr := []int{1, 100000}
	list := utils.ArrayToLinkedList(arr)
	pairSum(list)
}
