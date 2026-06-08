type MyQueue struct {
	s1 []int
	s2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		s1: []int{},
		s2: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.s2 = append(this.s1, x)
	this.s1 = this.s2
	this.s2 = []int{}
}

func (this *MyQueue) Pop() int {
	if len(this.s1) == 0 {
		return -1
	}
	val := this.s1[0]
	this.s1 = this.s1[1:]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.s1) == 0 {
		return -1
	}
	val := this.s1[0]
	return val
}

func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0
}
