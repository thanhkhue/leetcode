// https://leetcode.com/problems/count-primes/submissions/

package main

import "log"

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	primes := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		primes[i] = true
	}

	p := 2
	for p*p <= n {

		// p not changes since p is prime
		if primes[p] {

			// start from square of p
			// takes step of p to mark all are not
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}

		p++
	}

	log.Printf("Primes %+v", primes)
	cnt := 0
	for i := 0; i <= n; i++ {
		if primes[i] {
			log.Printf("Prime %d", i)
			cnt++
		}
	}

	return cnt
}

func main() {
	n := 50
	countPrimes(n)
}
