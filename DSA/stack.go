package main

import "fmt"

type Node[T any] struct {
	val  T
	next *Node[T]
}

type stack[T any] struct {
	head *Node[T]
	size int
}

func (s *stack[T]) push(newVal T) {
	newNode := &Node[T]{newVal, s.head}
	s.head = newNode
	s.size++
}

func (s *stack[T]) pop() *Node[T] {
	temp := s.head
	if s.head != nil {
		s.head = s.head.next
		s.size--
	}
	return temp
}


func (s *stack[T]) print(){
	temp:= s.head

	fmt.Print("[ ")
	for temp != nil {
		fmt.Print(temp.val , " ")
		temp = temp.next
	}
	fmt.Print("] \n")
}