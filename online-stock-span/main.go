/*
https://leetcode.com/problems/online-stock-span/
*/

package main

type Node struct {
	Val int
	Res int
}

type StockSpanner struct {
	Prices []Node
}

func Constructor() StockSpanner {
	return StockSpanner{
		Prices: make([]Node, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	res := 1
	if len(this.Prices) == 0 {
		this.Prices = append(this.Prices, Node{
			Val: price,
			Res: res,
		})
		return res
	}

	for len(this.Prices) > 0 && price >= this.Prices[len(this.Prices)-1].Val {
		res += this.Prices[len(this.Prices)-1].Res
		this.Prices = this.Prices[:len(this.Prices)-1]
	}
	this.Prices = append(this.Prices, Node{
		Val: price,
		Res: res,
	})
	return res
}
