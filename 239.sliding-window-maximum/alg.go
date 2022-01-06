package _39_sliding_window_maximum

type queue struct {
	list []int
	len  int
}

func (q *queue) push(v int) {
	q.list = append(q.list, v)
	q.len++
}

func (q *queue) popBack() int {
	if q.len == 0 {
		panic("empty")
	}
	q.len--
	rear := q.list[q.len]
	q.list = q.list[:q.len]
	return rear
}
func (q *queue) back() int {
	if q.len == 0 {
		panic("empty")
	}
	return q.list[q.len-1]
}

func (q *queue) popFront() int {
	if q.len == 0 {
		panic("empty")
	}
	q.len--
	front := q.list[0]
	q.list = q.list[1:]
	return front
}
func (q *queue) front() int {
	if q.len == 0 {
		panic("empty")
	}
	return q.list[0]
}
func maxSlidingWindow(nums []int, k int) []int {
	q := queue{}
	for i := 0; i < k; i++ {
		for q.len != 0 && nums[i] >= nums[q.back()] {
			q.popBack()
		}
		q.push(i)
	}
	length := len(nums)
	ans := make([]int, 1, length-k+1)
	ans[0] = nums[q.front()]
	for i := k; i < length; i++ {
		for q.len != 0 && nums[i] >= nums[q.back()] {
			q.popBack()
		}
		q.push(i)
		for q.front() <= i-k {
			q.popFront()
		}
		ans = append(ans, nums[q.front()])
	}

	return ans
}
