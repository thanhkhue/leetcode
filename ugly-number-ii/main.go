// https://leetcode.com/problems/ugly-number-ii/

package main

import (
	"log"
	"time"
)

func main() {
	n := 10
	now := time.Now()
	log.Printf("Result %d took %v", nthUglyNumber(n), time.Since(now))
}

func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	t2, t3, t5 := 0, 0, 0
	k := make([]int, n+1)
	k[0] = 1
	for i := 1; i <= n; i++ {
		t := min(min(k[t2]*2, k[t3]*3), k[t5]*5)
		log.Printf("t=%d", t)
		k[i] = t
		if k[i] == k[t2]*2 {
			t2++
		}
		if k[i] == k[t3]*3 {
			t3++
		}
		if k[i] == k[t5]*5 {
			t5++
		}
	}

	log.Printf("k[n-1]=%d, k[n]=%d, k[n+1]=%d", k[n-1], k[n], k[n])

	return 1

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func naive(n int) int {
	res := -1
	for i := 1; n > 0; i++ {
		if isUglyNumber(i) {
			n--
			res = i
		}
	}

	return res
}

func isUglyNumber(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}

	return n == 1
}
