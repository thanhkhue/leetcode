// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

package main

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GetTreeLevel(root *TreeNode) int {
	var level int
	getTreeLevel(root, &level)
	log.Printf("Level %+v", level)
	return level
}

func getTreeLevel(root *TreeNode, level *int) {
	if root == nil {
		return
	}

	(*level)++
	getTreeLevel(root.Left, level)
	getTreeLevel(root.Right, level)
}

func main() {

	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 8,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}

	/*

		root := &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 10,
			},
			Right: &TreeNode{
				Val: 4,
			},
		}

	*/

	log.Printf("Height of the tree is %d", findHeightOfTree(root))

	bfsVersion(root)
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}

func findHeightOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(findHeightOfTree(root.Left), findHeightOfTree(root.Right)) + 1
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return [][]int{{
			root.Val,
		}}
	}

	nodesWithLevels := make([][]int, 2000)
	level := 0
	genLevelOrderBottom(root, nodesWithLevels, level)

	res := make([][]int, 0)
	for i := len(nodesWithLevels) - 1; i >= 0; i-- {
		if len(nodesWithLevels[i]) == 0 {
			continue
		}

		res = append(res, nodesWithLevels[i])
	}

	log.Printf("res %+v", res)

	return nil
}

func genLevelOrderBottom(root *TreeNode, nodesWithLevels [][]int, level int) {
	if root == nil {
		return
	}
	nodesWithLevels[level] = append(nodesWithLevels[level], root.Val)
	level = level + 1
	genLevelOrderBottom(root.Left, nodesWithLevels, level)
	genLevelOrderBottom(root.Right, nodesWithLevels, level)
}

func bfsVersion(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		qLen := len(queue)
		nodesInCurrLevel := make([]int, 0)
		for i := 0; i < qLen; i++ {
			frontElem := queue[0]
			nodesInCurrLevel = append(nodesInCurrLevel, frontElem.Val)
			queue = queue[1:]
			if frontElem.Left != nil {
				queue = append(queue, frontElem.Left)
			}
			if frontElem.Right != nil {
				queue = append(queue, frontElem.Right)
			}
		}
		level++
		log.Printf("Elements at level %d are %+v", level, nodesInCurrLevel)
		res = append([][]int{nodesInCurrLevel}, res...)
	}

	log.Printf("Final results %+v", res)

	return res
}
