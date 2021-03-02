package main

//link:https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

type CQueue struct {
	In  []int
	Out []int
}

func Constructor() CQueue {
	return CQueue{
		In:  make([]int, 0),
		Out: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.In = append(this.In, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.Out) > 0 {
		outLen := len(this.Out)
		value := this.Out[outLen-1]
		this.Out = this.Out[0 : outLen-1]
		return value
	}
	if len(this.In) > 0 {
		inLen := len(this.In)
		for i := inLen - 1; i >= 0; i-- {
			this.Out = append(this.Out, this.In[i])
		}
		this.In = this.In[:0]
		outLen := len(this.Out)
		value := this.Out[outLen-1]
		this.Out = this.Out[0 : outLen-1]
		return value
	}
	return -1
}
