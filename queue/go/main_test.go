package main

import (
	"testing"
)

type impType string

const (
	enQueueImpType impType = "enqueue"
	deQueueImpType impType = "dequeue"
)

func TestQueue(t *testing.T) {
	testCase := []struct {
		name    string
		value   int
		impType impType
		want    error
	}{

		{name: "Dequeue0ErrEmpty", value: 0, impType: deQueueImpType, want: ErrQueueIsEmpty},

		{name: "Enqueue(1)", value: 1, impType: enQueueImpType, want: nil},
		{name: "Enqueue(2)", value: 2, impType: enQueueImpType, want: nil},
		{name: "Enqueue(3)", value: 3, impType: enQueueImpType, want: nil},
		{name: "Enqueue(4)", value: 4, impType: enQueueImpType, want: nil},
		{name: "Enqueue(5)", value: 5, impType: enQueueImpType, want: nil},
		{name: "Enqueue(6)ErrFull", value: 6, impType: enQueueImpType, want: ErrQueueIsFull},
		{name: "Dequeue1", value: 0, impType: deQueueImpType, want: nil},
		{name: "Enqueue(6)", value: 6, impType: enQueueImpType, want: nil},

		{name: "Dequeue2", impType: deQueueImpType, want: nil},
		{name: "Dequeue3", impType: deQueueImpType, want: nil},
		{name: "Dequeue4", impType: deQueueImpType, want: nil},
		{name: "Dequeue5", impType: deQueueImpType, want: nil},
		{name: "Dequeue6", impType: deQueueImpType, want: nil},
		{name: "Dequeue7ErrEmpty", impType: deQueueImpType, want: ErrQueueIsEmpty},
	}

	queue := NewQueue(5)

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			var got error
			switch v.impType {
			case deQueueImpType:
				got = queue.Dequeue()
			default:
				got = queue.Enqueue(v.value)
			}
			if got != v.want {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}
