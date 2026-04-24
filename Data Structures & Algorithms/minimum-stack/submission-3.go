type Node struct {
	val  int
	min  int
	next *Node
}

type MinStack struct {
	head *Node
}

func Constructor() MinStack {
	return MinStack{nil}
}

func (this *MinStack) Push(val int) {
	if this.head == nil {
		this.head = &Node{val: val, min: val}
		return
	}
	currentMin := this.head.min
	if val < currentMin {
		currentMin = val
	}

	newNode := &Node{val: val, min: currentMin, next: this.head}

	this.head = newNode
}

func (this *MinStack) Pop() {
	if this.head == nil {
		return
	}
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}
