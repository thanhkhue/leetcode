// https://leetcode.com/problems/rle-iterator/

package main

import "log"

type RLEIterator struct {
	encoding []int
	index    int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{
		encoding: encoding,
		index:    0,
	}
}

var remains = 0

func (this *RLEIterator) Next(n int) int {
	for this.index < len(this.encoding) && n > this.encoding[this.index] {
		this.index += 2
		n = n - this.encoding[this.index]
	}

	if this.index >= len(this.encoding) {
		return -1
	}

	this.encoding[this.index] = this.encoding[this.index] - n
	return this.encoding[this.index+1]
}

/*
// initially solution
func (this *RLEIterator) Next(n int) int {
    remains := 0
	for i := this.turn; i < len(this.encoding); i += 2 {
		if this.encoding[i] <= 0 {
			continue
		}
		n += remains
		if this.encoding[i] >= n {
			remains = 0
			this.encoding[i] -= n
			return this.encoding[i+1]
		}

        this.turn += 2
		remains = n - this.encoding[i]
		n = n - this.encoding[i]
	}

	return -1
}
*/

func main() {

	encoding := []int{811, 903, 310, 730, 899, 684, 472, 100, 434, 611}
	obj := Constructor(encoding)
	arr := []int{358, 345, 154, 265, 73, 220, 138, 4, 170, 88}
	for _, next := range arr {
		param := obj.Next(next)
		log.Printf("[Param]=%d", param)
	}

}
