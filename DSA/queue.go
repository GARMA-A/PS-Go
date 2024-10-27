package main

import "fmt"

type queue[T any] struct {
	front *Node[T]
	back  *Node[T]
	size  int
}

func (q *queue[T]) inqueue(newVal T) {
	newNode := &Node[T]{newVal, nil}
	if q.front == nil {
		q.front = newNode
		q.back = newNode
	} else {
		q.back.next = newNode
		q.back = newNode
	}
	q.size++
}

func (q *queue[T]) dequeue() {
	if q.front != nil {
		q.front = q.front.next
		q.size--
		if q.front == nil {
			q.back = nil

		}
	}
}

func (q *queue[T]) print(){
	temp := q.front
	fmt.Print("[ ")
	for temp != nil {
		fmt.Print(temp.val , " ")
		temp = temp.next
	}
	fmt.Print(" ]\n")
}
