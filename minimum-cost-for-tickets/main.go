// https://leetcode.com/problems/minimum-cost-for-tickets/

package main

import "log"

func main() {
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}
	log.Printf("Results %d", mincostTickets(days, costs))
}

func mincostTickets(days []int, costs []int) int {

	table := make([]int, days[len(days)-1]+1)
	uniqueDays := make(map[int]bool)
	for i := 0; i < len(days); i++ {
		uniqueDays[days[i]] = true
	}

	log.Printf("First day %d, last day %d", days[0], days[len(days)-1])

	for i := days[0]; i <= days[len(days)-1]; i++ {

		if !uniqueDays[i] {
			table[i] = table[i-1]
			continue
		}

		table[i] = min(min(
			table[max(0, (i-1))]+costs[0],
			table[max(0, (i-7))]+costs[1],
		),
			table[max(0, (i-30))]+costs[2],
		)

		log.Printf("1. %d", table[max(0, (i-1))]+costs[0])
		log.Printf("2. %d", table[max(0, (i-7))]+costs[1])
		log.Printf("3. %d", table[max(0, (i-30))]+costs[2])
		log.Printf("Choose %d", table[i])
	}

	log.Printf("table %+v", table)

	return table[days[len(days)-1]%30]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
