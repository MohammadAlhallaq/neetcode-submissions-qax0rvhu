type Node struct {
	val  int
	min  int
	next *Node
}

type MinStack struct {
	head *Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	node := &Node{
		min:  val,
		val:  val,
		next: this.head,
	}

	if this.head != nil && this.head.min < val {
		node.min = this.head.min
	}

	this.head = node
}

func (this *MinStack) Pop() {
	if this.head != nil {
		this.head = this.head.next
	}
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}
