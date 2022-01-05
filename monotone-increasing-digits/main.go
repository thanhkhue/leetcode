// https://leetcode.com/problems/monotone-increasing-digits/

package main

import (
	"log"
	"strconv"
)

func monotoneIncreasingDigits(n int) int {
	if isMonotoneIncrease(n) {
		return n
	}

	for n > 0 {
		n--
		if isMonotoneIncrease(n) {
			return n
		}
	}

	return 0
}

// Inefficient
func isMonotoneIncrease(n int) bool {
	str := strconv.Itoa(n)
	for i := 0; i < len(str)-1; i++ {
		if str[i] > str[i+1] {
			return false
		}
	}
	return true
}

func main() {
	n := 10
	log.Printf("Res %d", monotoneIncreasingDigits(n))
}
