/*
https://leetcode.com/problems/next-greater-element-i/description/
*/

package main

import "log"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreaterStacks := make([]int, 0)
	toNextGreater := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		for len(nextGreaterStacks) > 0 && nums2[i] > nextGreaterStacks[len(nextGreaterStacks)-1] {
			toNextGreater[nextGreaterStacks[len(nextGreaterStacks)-1]] = nums2[i]
			nextGreaterStacks = nextGreaterStacks[:len(nextGreaterStacks)-1]
		}
		nextGreaterStacks = append(nextGreaterStacks, nums2[i])
	}

	nextGreaterElements := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		nextGreater, ok := toNextGreater[nums1[i]]
		if !ok {
			nextGreaterElements[i] = -1
		} else {
			nextGreaterElements[i] = nextGreater
		}
	}

	return nextGreaterElements
}

func main() {

	nums1 := []int{1, 3, 5, 2, 4}
	nums2 := []int{6, 5, 4, 3, 2, 1, 7}
	log.Printf("nextGreaterElement %+v", nextGreaterElement(nums1, nums2))

	nums3 := []int{2, 4}
	nums4 := []int{1, 2, 3, 4}
	log.Printf("nextGreaterElement %+v", nextGreaterElement(nums3, nums4))
}
