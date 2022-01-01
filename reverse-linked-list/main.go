package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func newNode(n *Node, val int) *Node {
	if n == nil {
		return &Node{
			Val: val,
		}
	}

	if n.Next == nil {
		n.Next = &Node{
			Val: val,
		}
		return n
	}

	p := n
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &Node{Val: val}
	return n
}

func arrayToLinkedList(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	n := &Node{
		Val: arr[0],
	}
	for i := 1; i < len(arr); i++ {
		n = newNode(n, arr[i])
	}
	return n
}

func linkedListToArray(n *Node) []int {
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

func main() {
	arr := []int{5, 8, 12, 7, 2, 89, 333}
	n := arrayToLinkedList(arr)
	printLinkedList(n)
	naiveReverseLinkedList(n)
}

func reverseArray(arr []int) []int {
	r := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		r = append(r, arr[i])
	}
	return r
}

func reverseWithoutExtraMemory(n *Node) *Node {
	return nil
}

func naiveReverseLinkedList(n *Node) *Node {
	arr := linkedListToArray(n)
	newNodes := arrayToLinkedList(reverseArray(arr))
	fmt.Println("\nREVERSING")
	printLinkedList(newNodes)
	return nil
}

func printLinkedList(n *Node) {
	p := n
	for p.Next != nil {
		fmt.Printf("%d -> ", p.Val)
		p = p.Next
	}

	if p != nil {
		fmt.Printf("%d", p.Val)
	}
}
