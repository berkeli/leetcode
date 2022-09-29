package leetcode

type MyCircularQueue struct {
	data []int
	size int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		size: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.data) >= this.size {
		return false
	}
	this.data = append(this.data, value)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[1:]
	return true
}

func (this *MyCircularQueue) Front() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[0]
}

func (this *MyCircularQueue) Rear() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return len(this.data) == this.size
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
