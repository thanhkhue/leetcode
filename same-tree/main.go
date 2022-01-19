// https://leetcode.com/problems/same-tree/

package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	// qArr := []int{}
	// pArr := []int{}
	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	log.Printf("1.Is same tree %t", isSameTree(p, q))
	log.Printf("2.Is same tree %t", isSameTree(p, nil))
	log.Printf("3.Is same tree %t", isSameTree(nil, nil))
	log.Printf("4.Is same tree %t", isSameTree(nil, q))
}
