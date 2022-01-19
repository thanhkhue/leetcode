// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

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
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}
	connect(root)
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	levelToNodes := make(map[int][]*Node)
	level := 1
	bfsWithExtraMem(root, level, levelToNodes)
	for level, nodes := range levelToNodes {
		if len(nodes) > 0 {
			log.Printf("Level %d", level)
			for i := 0; i < len(nodes); i++ {
				if nodes[i].Next == nil {
					fmt.Printf("Next of %d is NULL", nodes[i].Val)
				} else {
					fmt.Printf("Next of %d -> %d\n", nodes[i].Val, nodes[i].Next.Val)
				}

				// fmt.Printf("%d -> ", nodes[i].Val)
			}
			fmt.Println()
		}
		// log.Printf("Level %d has %d nodes, first node %+v", level, len(nodes), nodes[0])
	}

	return root
}

func bfsWithExtraMem(root *Node, level int, levelToNodes map[int][]*Node) {
	if root == nil {
		return
	}

	// if levelToNodes[level] == nil {
	// 	levelToNodes[level] = make([]*Node, 0)
	// }
	nodes := levelToNodes[level]
	if len(nodes) > 0 {
		nodes[len(nodes)-1].Next = root
	}
	levelToNodes[level] = append(levelToNodes[level], root)
	level++
	bfsWithExtraMem(root.Left, level, levelToNodes)
	bfsWithExtraMem(root.Right, level, levelToNodes)
}

func bfs(root *Node) {
	if root == nil {
		return
	}

	if root.Left != nil {
		root.Left.Next = root.Right
	}

	if root.Next != nil && root.Right != nil {
		root.Right.Next = root.Next.Left
	}

	bfs(root.Left)
	bfs(root.Right)
}
