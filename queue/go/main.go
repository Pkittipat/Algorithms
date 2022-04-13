package main

import (
	"errors"
	"fmt"
)

type Queue struct {
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

func NewQueue(length int) *Queue {
	return &Queue{
		Length: length,
		Queue:  make([]*int, length),
		Head:   InitValue,
		Tail:   InitValue,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.Tail == q.Head
}

func (q *Queue) IsFull(tail int) bool {
	if tail < q.Head {
		return (q.Head-tail)%q.Length <= 0
	}
	return (tail-q.Head)%q.Length <= 0
}

func (q *Queue) Enqueue(val int) error {
	switch {
	case q.Tail == -1:
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

func (q *Queue) Dequeue() error {
	if q.Head == -1 {
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

func main() {
	queue := NewQueue(5)
	fmt.Println(queue.Queue)
	testCase := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, v := range testCase {
		if v == 3 {
			err := queue.Dequeue()
			if err != nil {
				break
			}
		} else {
			err := queue.Enqueue(v)
			if err != nil {
				break
			}
		}
	}
	fmt.Println(queue.Queue)
}
