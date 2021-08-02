package arrayqueue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zergon321/arrayqueue"
)

func TestQueueLength(t *testing.T) {
	queue, err := arrayqueue.NewQueue(512)
	assert.Nil(t, err)

	for i := 0; i < 512; i++ {
		queue.Enqueue(i)
	}

	assert.Equal(t, 512, queue.Length())

	for i := 0; i < 512; i++ {
		_, err := queue.Dequeue()
		assert.Nil(t, err)
	}

	assert.Equal(t, 0, queue.Length())
}

func TestQueueFIFO(t *testing.T) {
	queue, err := arrayqueue.NewQueue(512)
	assert.Nil(t, err)

	queue.Enqueue(0)
	queue.Enqueue(1)

	el, err := queue.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 0, el)

	el, err = queue.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 1, el)
}

func TestQueueBufferSize(t *testing.T) {
	queue, err := arrayqueue.NewQueue(512)
	assert.Nil(t, err)

	assert.Equal(t, 0, queue.Length())
	assert.Equal(t, 512, queue.BufferSize())

	err = queue.SetBufferSize(256)
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Length())
	assert.Equal(t, 256, queue.BufferSize())

	queue.Enqueue(2)
	err = queue.SetBufferSize(128)
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Length())
	assert.Equal(t, 128, queue.BufferSize())
}

func TestQueueClear(t *testing.T) {
	queue, err := arrayqueue.NewQueue(512)
	assert.Nil(t, err)

	for i := 0; i < 512; i++ {
		queue.Enqueue(i)
	}

	queue.Clear()
	assert.Equal(t, 0, queue.Length())
	assert.Equal(t, 512, queue.BufferSize())
}
