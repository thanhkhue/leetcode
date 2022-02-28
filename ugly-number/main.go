// https://leetcode.com/problems/ugly-number/

/*
	A number is said to be ugly if it has only 2, 3 and 5
	as its prime factors.
	In other words, if a number can only obtained by multiplying
	power for 2, 3 and 5 the number is said to be an ugly number
*/

package main

import "log"

func findFactors(n int) map[int]bool {
	f := make(map[int]bool)
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			f[i] = true
		}
	}
	return f
}

func isUgly(n int) bool {

	if n <= 0 {
		return false
	}

	for n%2 == 0 {
		n = n / 2
	}

	for n%3 == 0 {
		n = n / 3
	}

	for n%5 == 0 {
		n = n / 5
	}

	// only divide for 2, 3, 5
	// then result must be 1
	return n == 1
}

func main() {
	for i := 1; i <= 15; i++ {
		if isUgly(i) {
			log.Printf("%d is an ugly number", i)
		} else {
			log.Printf("%d is not an ugly number", i)
		}
	}
}

func allPrimesLessThanN(n int) []int {

	primes := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		primes[i] = true
	}
	p := 2
	for p*p <= n {
		if primes[p] {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
		p++
	}

	res := make([]int, 0)
	for i := 3; i <= n; i++ {
		if primes[i] {
			res = append(res, i)
		}
	}

	return res
}
