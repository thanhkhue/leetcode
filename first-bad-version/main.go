package main

import (
	"log"
	"math"
)

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {

	if n <= 1 {
		return n
	}

	start, end := 0, n
	badVersion := math.MaxInt
	log.Printf("badVersion %d", badVersion)
	for start <= end {
		mid := (start + end) / 2
		if isBadVersion(mid) {
			if mid <= badVersion {
				badVersion = mid
			}
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	log.Printf("badVersion %d", badVersion)

	return n
}

func isBadVersion(version int) bool {
	return version >= BadVersion
}

var BadVersion = 4

func main() {

	firstBadVersion(5)
}
