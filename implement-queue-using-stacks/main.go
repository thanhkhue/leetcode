// https://leetcode.com/problems/implement-queue-using-stacks/

package main

import "log"

type MyQueue struct {
	PushStack, PopStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		PushStack: make([]int, 0),
		PopStack:  make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.PushStack = append(this.PushStack, x)
}

func (this *MyQueue) Peek() int {
	if len(this.PopStack) == 0 {
		for len(this.PushStack) > 0 {
			this.PopStack = append(this.PopStack, this.PushStack[len(this.PushStack)-1])
			this.PushStack = this.PushStack[:len(this.PushStack)-1]
		}
	}
	return this.PopStack[len(this.PopStack)-1]
}

func (this *MyQueue) Pop() int {
	popped := this.Peek()
	this.PopStack = this.PopStack[:len(this.PopStack)-1]
	return popped
}

func (this *MyQueue) Empty() bool {
	return len(this.PushStack) == 0 && len(this.PopStack) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	peeked1 := obj.Peek()
	log.Printf("Peek value %d, push stack %+v, pop stack %+v", peeked1, obj.PushStack, obj.PopStack)
	peeked2 := obj.Peek()
	log.Printf("Peek value %d, push stack %+v, pop stack %+v", peeked2, obj.PushStack, obj.PopStack)

	popped := obj.Pop()
	log.Printf("Popped value %d, push stack %+v, pop stack %+v", popped, obj.PushStack, obj.PopStack)

	popped = obj.Pop()
	log.Printf("Popped value %d, push stack %+v, pop stack %+v", popped, obj.PushStack, obj.PopStack)

	log.Printf("is stack empty? %t", obj.Empty())
	obj.Pop()
	log.Printf("is stack empty? %t", obj.Empty())
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
