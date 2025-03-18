package queues

// Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).
//
// Implement the MyQueue class:
//
//	void push(int x) Pushes element x to the back of the queue.
//	int pop() Removes the element from the front of the queue and returns it.
//	int peek() Returns the element at the front of the queue.
//	boolean empty() Returns true if the queue is empty, false otherwise.
//
// Notes:
//
//	You must use only standard operations of a stack, which means only push to top, peek/pop from top, size, and is empty operations are valid.
//	Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.
type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

func (q *MyQueue) Pop() int {
	if !q.Empty() {
		last := q.Peek()
		q.out = q.out[:len(q.out)-1]
		return last
	}
	return 0
}

func (q *MyQueue) Peek() int {
	if len(q.out) == 0 {
		for i := len(q.in) - 1; i >= 0; i-- {
			q.out = append(q.out, q.in[i])
		}
		q.in = make([]int, 0)
	}
	if len(q.out) > 0 {
		return q.out[len(q.out)-1]
	}
	return 0
}

func (q *MyQueue) Empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}
