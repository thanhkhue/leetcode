/*
https://leetcode.com/problems/merge-two-sorted-lists/
*/
package main

import (
	"fmt"
	"log"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var Method = "newLinkedList"

func main() {
	arr1 := []int{1}
	arr2 := []int{2}
	result := mergeTwoLists(arrayToLinkedList(arr1), arrayToLinkedList(arr2))
	fmt.Printf("Result sorted %+v", linkedListToArray(result))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if Method == "array" {
		return usingArray(list1, list2)
	}
	if Method == "newLinkedList" {
		return usingNewLinkedList(list1, list2)
	}
	return nil
}

func usingNewLinkedList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	list := new(ListNode)
	head := list // store head
	for list1 != nil || list2 != nil {
		if list1.Val > list2.Val {
			log.Printf("list2.Val %d", list2.Val)
			list.Val = list2.Val
			list2 = list2.Next
		} else if list1.Val == list2.Val {
			log.Printf("list1.Val %+v", list1)

			list.Val = list1.Val
			list.Next = new(ListNode)
			list = list.Next
			list.Val = list2.Val
			list1 = list1.Next
			list2 = list2.Next
		} else {
			log.Printf("list1.Val %d", list1.Val)
			list.Val = list1.Val
			list1 = list1.Next
		}
		list.Next = new(ListNode)
		list = list.Next
	}

	if list1 != nil && list2 != nil {
		list = list.Next
		list.Next = new(ListNode)
		if list1.Val > list2.Val {
			list.Val = list2.Val
			list.Next.Val = list1.Val
		} else {
			list.Val = list1.Val
			list.Next.Val = list2.Val
		}
	} else if list1 != nil {
		list = list.Next
		list.Next = new(ListNode)
		list.Val = list1.Val
	} else if list2 != nil {
		list = list.Next
		if list == nil {
			list = new(ListNode)
		}
		list.Val = list2.Val
	}

	return head
}

func usingArray(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	arr := append(linkedListToArray(list1), linkedListToArray(list2)...)
	sort.Ints(arr)
	return arrayToLinkedList(arr)
}

func linkedListToArray(l *ListNode) []int {
	arr := make([]int, 0)
	for l.Next != nil {
		arr = append(arr, l.Val)
		l = l.Next
	}
	return arr
}

func arrayToLinkedList(arr []int) *ListNode {
	l := new(ListNode)
	head := l
	for i := 0; i < len(arr); i++ {
		head.Val = arr[i]
		if head.Next == nil {
			head.Next = new(ListNode)
		}
		head = head.Next
	}
	return l
}
