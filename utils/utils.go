package utils

type ListNode struct {
	Val  int
	Next *ListNode
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

func ArrayToLinkedList(arr []int) *ListNode {
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

func LinkedListToArray(n *ListNode) []int {
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
