package main

//link:https://leetcode-cn.com/problems/implement-queue-using-stacks/
type MyQueue struct {
	first  *Stack
	second *Stack
}

type Stack struct {
	Element []int
	count   int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		first: &Stack{
			Element: make([]int, 0),
			count:   0,
		},
		second: &Stack{
			Element: make([]int, 0),
			count:   0,
		},
	}
}

func (s *Stack) Empty() bool {
	return s.count == 0
}
func (s *Stack) Top() int {
	var val int
	if !s.Empty() {
		return s.Element[s.count-1]
	}
	return val
}

func (s *Stack) Pop() int {
	var val int
	if !s.Empty() {
		val = s.Element[s.count-1]
		s.Element = s.Element[:s.count-1]
		s.count--
	}
	return val
}
func (s *Stack) Push(x int) {
	if s.Element != nil {
		s.Element = append(s.Element, x)
		s.count++
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.first.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.second.Empty() {
		for !this.first.Empty() {
			this.second.Push(this.first.Pop())
		}
	}
	return this.second.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.second.Empty() {
		for !this.first.Empty() {
			this.second.Push(this.first.Pop())
		}
	}
	return this.second.Top()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.first.Empty() && this.second.Empty()
}
