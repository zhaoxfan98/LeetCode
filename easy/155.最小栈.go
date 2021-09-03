package easy

import "math"

type MinStack struct {
	items    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		items:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(top, val))
}

func (this *MinStack) Pop() {
	this.items = this.items[:len(this.items)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
