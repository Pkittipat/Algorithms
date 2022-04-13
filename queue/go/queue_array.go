package queue

import (
	"errors"
)

type QueueArray struct {
	Length int
	Queue  []*int
	Head   int
	Tail   int
}

const (
	InitValue = -1
)

var (
	ErrQueueIsFull  = errors.New("queue is fully")
	ErrQueueIsEmpty = errors.New("queue is empty")
)

func NewQueue(length int) *QueueArray {
	return &QueueArray{
		Length: length,
		Queue:  make([]*int, length),
		Head:   InitValue,
		Tail:   InitValue,
	}
}

func (q *QueueArray) IsEmpty() bool {
	return q.Tail == q.Head
}

func (q *QueueArray) IsFull(tail int) bool {
	if tail < q.Head {
		return (q.Head-tail)%q.Length <= 0
	}
	return (tail-q.Head)%q.Length <= 0
}

func (q *QueueArray) Enqueue(val int) error {
	switch {
	case q.Tail == InitValue:
		q.Head = 0
		q.Tail = 0
		q.Queue[q.Tail] = &val
	default:
		buf := q.Tail + 1
		buf %= q.Length

		if q.IsFull(buf) {
			return ErrQueueIsFull
		}
		q.Tail = buf
		q.Queue[q.Tail] = &val
	}
	return nil
}

func (q *QueueArray) Dequeue() error {
	if q.Head == InitValue {
		return ErrQueueIsEmpty
	}

	if q.IsEmpty() {
		if q.Queue[q.Head] == nil {
			return ErrQueueIsEmpty
		}
		q.Queue[q.Head] = nil
		return nil
	}

	q.Queue[q.Head] = nil
	q.Head++
	q.Head %= q.Length
	return nil
}
