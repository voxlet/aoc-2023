package priorityqueue

func (pq PriorityQueue[T]) Len() int {
	return len(pq.Entries)
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq.IsBefore(pq.Entries[i].Priority, pq.Entries[j].Priority)
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq.Entries[i], pq.Entries[j] = pq.Entries[j], pq.Entries[i]
	pq.Entries[i].index = i
	pq.Entries[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	entry := x.(Entry[T])
	entry.index = len(pq.Entries)
	pq.Entries = append(pq.Entries, entry)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := pq.Entries
	n := len(old)
	entry := old[n-1]

	old[n-1].Value = nil // avoid memory leak
	entry.index = -1     // for safety

	pq.Entries = old[0 : n-1]

	return entry
}
