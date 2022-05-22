package main

import (
	priorityQ "datastructures/PriorityQueue"
	"fmt"
)

func main() {
	pq := priorityQ.NewPriorityQ[int]()
	pq.Insert(1)
	pq.Insert(100)
	pq.Insert(9)
	pq.Insert(0)
	pq.Insert(-23)

	for !pq.IsEmpty() {
		if top, err := pq.Top(); err == nil {
			if err := pq.Pop(); err != nil {
				return
			}
			fmt.Println(top)
		}
	}
}
