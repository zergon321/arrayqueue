package arrayqueue

// Queue is a queue based on a slice.
type Queue struct {
	data []interface{}
}

// Length returns the numbers of elements
// in the queue.
func (queue *Queue) Length() int {
	return len(queue.data)
}

// Peek returns the last element of
// the queue.
func (queue *Queue) Peek() interface{} {
	lastIdx := len(queue.data) - 1
	lastElem := queue.data[lastIdx]

	return lastElem
}

// Enqueue puts the element in the end of the queue.
func (queue *Queue) Enqueue(elem interface{}) {
	queue.data = append([]interface{}{elem}, queue.data...)
}

// Dequeue returns the last element of the queue
// and removes it from the queue.
func (queue *Queue) Dequeue() interface{} {
	lastIdx := len(queue.data) - 1
	lastElem := queue.data[lastIdx]

	queue.data = append(queue.data[:lastIdx])

	return lastElem
}

// NewQueue returns a new queue
// based on a slice.
func NewQueue() *Queue {
	return &Queue{
		data: []interface{}{},
	}
}
