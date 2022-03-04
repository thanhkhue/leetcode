// https://leetcode.com/problems/majority-element/

package main

import (
	"log"
	"math"
)

func main() {
	arr := []int{3, 2, 3, 3, 2, 2, 2}
	log.Printf("Res %d", daq(arr, 0, len(arr)-1))
}

func daq(nums []int, low, high int) int {
	if low == high {
		return nums[low]
	}

	mid := (low + high) / 2
	left := daq(nums, low, mid)
	right := daq(nums, mid+1, high)

	if left == right {
		return left
	}

	leftOccurrencies := countOccurrency(nums, low, mid, left)
	rightOccurrencies := countOccurrency(nums, mid+1, high, right)
	if leftOccurrencies > rightOccurrencies {
		return left
	}

	return right
}

func countOccurrency(nums []int, low, high, target int) int {
	counter := 0
	for i := low; i <= high; i++ {
		if nums[i] == target {
			counter++
		}
	}
	return counter
}

func mooreVoting(nums []int) int {
	counter := 0
	m := math.MaxInt
	for _, v := range nums {
		if counter == 0 {
			m = v
			counter = 1
		} else if m == v {
			counter++
		} else {
			counter--
		}
	}
	log.Printf("Counter %d, m = %d", counter, m)
	return m
}

func majorityElement(nums []int) int {
	majorSize := len(nums) / 2
	greatest := math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] > greatest {
			greatest = nums[i]
		}
	}

	occurrencies := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		occurrencies[nums[i]]++
	}

	majorEle := 0
	maxOccur := -1
	for val, appreance := range occurrencies {
		if appreance > majorSize && appreance > maxOccur {
			maxOccur = appreance
			majorEle = val
		}
	}

	log.Printf("Major element %d, with occ %d", majorEle, maxOccur)

	return majorEle
}
