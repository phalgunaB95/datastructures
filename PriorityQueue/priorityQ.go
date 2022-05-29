package priorityQ

import (
	"errors"
	"golang.org/x/exp/constraints"
)

type PriorityQ[K constraints.Ordered] interface {
	Insert(key K)
	Pop() error
	Top() (K, error)
	IsEmpty() bool
	Len() int
	heapify(i int)
}

type pqHeap[K constraints.Ordered] []K

func (pq *pqHeap[K]) Insert(key K) {
	i := len(*pq)
	*pq = append(*pq, key)
	for i > 0 && (*pq)[i/2] > key {
		(*pq)[i] = (*pq)[i/2]
		i /= 2
	}
	(*pq)[i] = key
}
func (pq *pqHeap[K]) Pop() error {
	if len(*pq) == 0 {
		return errors.New("Heap is empty!")
	}
	(*pq)[0] = (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	pq.heapify(0)
	return nil
}
func (pq *pqHeap[K]) Top() (K, error) {
	var result K
	if len(*pq) == 0 {
		return result, errors.New("Heap is empty!")
	}
	result = (*pq)[0]
	return result, nil
}
func (pq *pqHeap[K]) IsEmpty() bool {
	return len(*pq) == 0
}
func (pq *pqHeap[K]) heapify(i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < len(*pq) && (*pq)[left] < (*pq)[smallest] {
		smallest = left
	}
	if right < len(*pq) && (*pq)[right] < (*pq)[smallest] {
		smallest = right
	}
	if smallest != i {
		(*pq)[smallest], (*pq)[i] = (*pq)[i], (*pq)[smallest]
		pq.heapify(smallest)
	}
}
func (pq *pqHeap[K]) Len() int {
	return len(*pq)
}
func NewPriorityQ[K constraints.Ordered]() PriorityQ[K] {
	tmpSlice := make([]K, 0)
	tmpHeap := pqHeap[K](tmpSlice)
	return &tmpHeap
}
