package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{1, 2}
	n := 1
	head := arrayToLinkedList(arr)
	result := solution2(head, n)
	log.Printf("Results %+v", linkedListToArray(result))
}

func solution2(head *ListNode, n int) *ListNode {
	dummy := head
	counter := 0
	for head != nil && counter <= n {
		head = head.Next
		counter++
	}

	if head == nil {
		return dummy.Next
	}

	slow, fast := dummy, head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	if slow == nil {
		dummy.Next = nil
		return dummy
	}

	slow.Next = slow.Next.Next
	return dummy
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil && n == 1 {
		return nil
	}
	last := head
	counter1 := 0
	for last != nil {
		counter1++
		last = last.Next
	}

	if n >= counter1 {
		return head.Next
	}

	log.Printf("counter1 %d, n=%d", counter1, n)

	dummy := head
	counter2 := 0
	for head != nil {
		counter2++
		if counter2 == counter1-n {
			log.Println("hello")
			if head.Next != nil {
				next := head.Next
				head.Next = next.Next
			} else {
				log.Println("hihihi")
				break
			}
		} else {
			head = head.Next
		}
	}

	return dummy
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

func linkedListToArray(n *ListNode) []int {
	if n == nil {
		return []int{}
	}
	arr := make([]int, 0)
	for n.Next != nil {
		arr = append(arr, n.Val)
		n = n.Next
	}

	return append(arr, n.Val)
}
