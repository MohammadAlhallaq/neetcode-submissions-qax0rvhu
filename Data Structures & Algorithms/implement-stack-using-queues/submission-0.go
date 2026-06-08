type MyStack struct {
	stack []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() int {
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return val
}

func (this *MyStack) Top() int {
	val := this.stack[len(this.stack)-1]
	return val
}

func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}