// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/

package main

import (
	"fmt"
	"log"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {
	/*
		root := &Node{
			Val: 1,
			Left: &Node{
				Val: 2,
				Left: &Node{
					Val: 4,
				},
				Right: &Node{
					Val: 5,
				},
			},
			Right: &Node{
				Val: 3,
				Right: &Node{
					Val: 7,
				},
			},
		}
	*/

	root := &Node{
		Val: 2,
		Left: &Node{
			Val: 1,
			Left: &Node{
				Val: 0,
				Left: &Node{
					Val: 2,
				},
			},
			Right: &Node{
				Val: 7,
				Left: &Node{
					Val: 1,
				},
				Right: &Node{
					Val: 0,
					Left: &Node{
						Val: 7,
					},
				},
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 9,
			},
			Right: &Node{
				Val: 1,
				Left: &Node{
					Val: 8,
				},
				Right: &Node{
					Val: 8,
				},
			},
		},
	}
	root = connect(root)
	printNextNode(root)
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	if root.Right != nil {
		nextOfRoot := root.Next
		for nextOfRoot != nil {
			if nextOfRoot.Left != nil {
				root.Right.Next = nextOfRoot.Left
				break
			}
			if nextOfRoot.Right != nil {
				root.Right.Next = nextOfRoot.Right
				break
			}
			nextOfRoot = nextOfRoot.Next
		}
	}

	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			nextOfRoot := root.Next
			for nextOfRoot != nil {
				if nextOfRoot.Left != nil {
					root.Left.Next = nextOfRoot.Left
					break
				}
				if nextOfRoot.Right != nil {
					root.Left.Next = nextOfRoot.Right
					break
				}
				nextOfRoot = nextOfRoot.Next
			}
		}
	}

	if root != nil && root.Right != nil && root.Right.Val == 0 {
		log.Printf("Root val %d, next of root %d, next of next %d", root.Val, root.Next.Val, root.Next.Next.Val)
	}

	connect(root.Right)
	connect(root.Left)

	return root
}

func printNextNode(root *Node) {
	if root == nil {
		return
	}

	if root.Next == nil {
		fmt.Printf("Next of %d is nil\n", root.Val)
	} else {
		fmt.Printf("Next of %d is %d\n", root.Val, root.Next.Val)
	}
	printNextNode(root.Left)
	printNextNode(root.Right)
}

func bfsWithExtraMem(root *Node, level int, levelToNodes map[int][]*Node) {
	if root == nil {
		return
	}

	nodes := levelToNodes[level]
	if len(nodes) > 0 {
		nodes[len(nodes)-1].Next = root
	}
	levelToNodes[level] = append(levelToNodes[level], root)
	level++
	bfsWithExtraMem(root.Left, level, levelToNodes)
	bfsWithExtraMem(root.Right, level, levelToNodes)
}
