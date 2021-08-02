package arrayqueue

import "fmt"

// Queue is a queue based on a slice.
type Queue struct {
	data []interface{}
}

// BufferSize returns the capacity of the queue.
func (queue *Queue) BufferSize() int {
	return cap(queue.data)
}

// SetBufferSize sets the capacity of the queue.
func (queue *Queue) SetBufferSize(bufferSize int) error {
	if bufferSize < 0 {
		return fmt.Errorf(
			"buffer size is less than 0: %d", bufferSize)
	}

	if bufferSize < len(queue.data) {
		return fmt.Errorf(
			"buffer size is less than the length of the queue: %d",
			bufferSize)
	}

	data := make([]interface{}, len(queue.data), bufferSize)

	copy(data, queue.data)
	queue.data = data

	return nil
}

// Clear removes all the elements from the queue.
func (queue *Queue) Clear() {
	queue.data = queue.data[:0]
}

// Length returns the numbers of elements
// in the queue.
func (queue *Queue) Length() int {
	return len(queue.data)
}

// Peek returns the last element of
// the queue.
func (queue *Queue) Peek() interface{} {
	return queue.data[0]
}

// Enqueue puts the element in the end of the queue.
func (queue *Queue) Enqueue(elem interface{}) {
	queue.data = append(queue.data, elem)
}

// Dequeue returns the last element of the queue
// and removes it from the queue.
func (queue *Queue) Dequeue() interface{} {
	lastElem := queue.data[0]
	queue.data = queue.data[1:]

	return lastElem
}

// NewQueue returns a new queue
// based on a slice.
func NewQueue(queueCapacity int) *Queue {
	data := make([]interface{}, 0, queueCapacity)

	return &Queue{
		data: data,
	}
}
