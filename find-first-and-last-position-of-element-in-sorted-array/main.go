// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

package main

import (
	"log"
)

func searchRange(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// mid := nums[len(nums)/2]

	return []int{}
}

func binarySearchV2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	numsLen := len(nums)
	m := -1
	firstInd, lastInd := -1, -1
	left, right := 0, numsLen-1
	for left <= right {
		m = (left + right) / 2
		if nums[m] < target {
			left = m + 1
		} else if nums[m] > target {
			right = m - 1
		} else {
			log.Printf("Reach target index %d, value %d", m, nums[m])
			break
		}
	}

	log.Println("Ever reach here")
	log.Printf("m %d", m)

	if nums[m] != target {
		log.Println("Return from here")
		log.Printf("first index %d, last index %d", firstInd, lastInd)
		return []int{firstInd, lastInd}
	}

	log.Printf("Found mid index %d", m)
	firstInd, lastInd = m, m

	findFirst, findLast := false, false
	for i := m; i < numsLen; i++ {
		if nums[i] != target && !findLast {
			lastInd = i - 1
			findLast = true
		}
	}
	for i := m; i > 0; i-- {
		if nums[i] != target && !findFirst {
			firstInd = i + 1
			findFirst = true
		}
	}

	log.Printf("1.findFirst %d, findLast %d", firstInd, lastInd)

	if firstInd == -1 {
		firstInd = 0
	}

	if lastInd == -1 {
		if firstInd > -1 {
			lastInd = firstInd
		} else {
			lastInd = len(nums) - 1
		}
	}
	log.Printf("2.findFirst %d, findLast %d", firstInd, lastInd)

	return []int{firstInd, lastInd}
}

func naive(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	firstInd, lastInd := -1, -1
	findFirst := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			if !findFirst {
				firstInd, lastInd = i, i
			}
			findFirst = true
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != target {
					break
				}
				lastInd = j
			}
		}
	}

	log.Printf("first index %d, last index %d of %d", firstInd, lastInd, target)

	return []int{firstInd, lastInd}
}

func main() {

	nums := []int{5, 7, 7, 8, 8, 9, 10, 12, 14, 16, 16}
	target := 7
	// nums := []int{2, 2}
	// target := 2
	// nums := []int{1, 3}
	// target := 1
	binarySearch(nums, target)

}

func binarySearch(nums []int, target int) []int {
	res := make([]int, 2)
	res[0] = findFirst(nums, target)
	res[1] = findLast(nums, target)
	log.Printf("Res %+v", res)
	return res
}

func findLast(nums []int, target int) int {
	start, end, idx := 0, len(nums)-1, -1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
		if nums[mid] == target {
			idx = mid
		}
	}
	log.Printf("Final index of first %d", idx)
	return idx
}

func findFirst(nums []int, target int) int {
	start, end, idx := 0, len(nums)-1, -1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] >= target {
			end = mid - 1
		} else {
			start = mid + 1
		}
		if nums[mid] == target {
			idx = mid
		}
	}
	log.Printf("Final index of last %d", idx)
	return idx
}
