package priorityqueue

import (
	"container/heap"
)

// An Entry is something we manage in a priority queue.
type Entry[T any] struct {
	Value    *T  // The value of the item; arbitrary.
	Priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T any] struct {
	IsBefore func(priorityA, priorityB int) bool
	Entries  []Entry[T]
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue[T]) Update(item *Entry[T], priority int) {
	item.Priority = priority
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue[T]) InsertValue(value *T, priority int) {
	heap.Push(pq, Entry[T]{
		Value:    value,
		Priority: priority,
	})
}

func (pq *PriorityQueue[T]) PopEntry() Entry[T] {
	return heap.Pop(pq).(Entry[T])
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.Entries) == 0
}
