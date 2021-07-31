package arrayqueue

import "fmt"

// Queue is a queue based on a slice.
type Queue struct {
	start int
	end   int
	data  []interface{}
}

func (queue *Queue) Empty() bool {
	return queue.start == queue.end &&
		queue.data[queue.start] == nil
}

func (queue *Queue) Count() int {
	if queue.start == queue.end {
		if queue.data[queue.start] == nil {
			return 0
		}

		return 1
	}

	return queue.end - queue.start + 1
}

func (queue *Queue) moveToEnd() {
	newEnd := len(queue.data) - 1
	distance := newEnd - queue.end

	for i := queue.end; i >= queue.start; i-- {
		queue.data[i+distance] = queue.data[i]
		queue.data[i] = nil
	}

	queue.end = newEnd
	queue.start += distance
}

func (queue *Queue) Enqueue(elem interface{}) error {
	if elem == nil {
		return fmt.Errorf("the element is nil")
	}

	// If the queue is empty.
	if queue.Empty() {
		queue.data[queue.start] = elem
		return nil
	}

	// If the queue is full.
	if queue.start == 0 && queue.end == len(queue.data)-1 {
		tmp := []interface{}{elem}

		tmp = append(tmp, queue.data...)
		queue.data = tmp
		queue.data = queue.data[:cap(queue.data)]
		queue.moveToEnd()

		return nil
	}

	// If there's some space on the right side.
	if queue.start == 0 && queue.end < len(queue.data)-1 {
		queue.moveToEnd()
	}

	queue.start--
	queue.data[queue.start] = elem

	return nil
}

func (queue *Queue) Dequeue() (interface{}, error) {
	// If the queue is empty.
	if queue.Empty() {
		return nil, fmt.Errorf(
			"the queue is empty, nothing to dequeue")
	}

	elem := queue.data[queue.end]
	queue.data[queue.end] = nil

	if queue.Count() > 1 {
		queue.end--
	}

	return elem, nil
}

func NewQueue(buffSize int) *Queue {
	return &Queue{
		start: 0,
		end:   0,
		data:  make([]interface{}, buffSize, buffSize),
	}
}
