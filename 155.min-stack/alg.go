package _55_min_stack

type MinStack struct {
	stack    []int
	minValue int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minValue: -1,
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
		this.minValue = val
	} else {
		diff := val - this.minValue
		this.stack = append(this.stack, diff)
		this.minValue = min(this.minValue, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		diff := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		if diff < 0 {
			this.minValue = this.minValue - diff
		}
	}
}

func (this *MinStack) Top() int {
	diff := this.stack[len(this.stack)-1]
	if diff < 0 {
		return this.minValue
	}
	return diff + this.minValue
}

func (this *MinStack) GetMin() int {
	return this.minValue
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
