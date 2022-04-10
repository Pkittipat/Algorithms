package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Length int
	Queue  []int
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
		Queue:  make([]int, length),
		Head:   InitValue,
		Tail:   InitValue,
	}
}

// [6, 0, 0, 0, 0]
// tail: -1, 0, 1, 2, 3, 4, 0
// head: -1, 0, 1, 2, 3, 4, 0

func (q *Queue) IsEmpty() bool {
	fmt.Printf("tail %v head %v\n", q.Tail, q.Head)
	return q.Tail == q.Head
}

func (q *Queue) IsFull() bool {
	if q.Head == InitValue {
		return false
	}
	return (q.Head - q.Tail) == 1
}

func (q *Queue) Enqueue(val int) error {
	if q.IsFull() {
		return ErrQueueIsFull
	}
	if q.Head == InitValue && q.Tail == InitValue {
		q.Head += 1
		q.Tail += 1
	}
	nextIndex := (q.Tail + 1) % q.Length
	q.Queue[nextIndex] = val
	if nextIndex == q.Head {
		nextIndex -= 1
	}
	q.Tail = nextIndex

	return nil
}

func (q *Queue) Dequeue() error {
	if q.IsEmpty() {
		q.Queue[q.Head] = 0
		return ErrQueueIsEmpty
	}
	nextIndex := (q.Head + 1) % q.Length
	q.Queue[nextIndex] = 0
	q.Head = nextIndex
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
