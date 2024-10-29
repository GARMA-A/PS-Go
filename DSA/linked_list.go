package main

import "fmt"


type linked_list[T any] struct {
	head *Node[T]
	size int
}

// O(1)
func (ls *linked_list[T]) addFront(newVal T) {
	newNode := &Node[T]{newVal, ls.head}
	ls.head = newNode
	ls.size++
}

// O(n)
func (ls *linked_list[T]) addBack(newVal T) {
	newNode := &Node[T]{newVal, nil}
	if ls.head != nil {
		tempNode := ls.head
		for tempNode.next != nil {
			tempNode = tempNode.next
		}
		tempNode.next = newNode
	} else {
		ls.head = newNode
	}

	ls.size++
}

// O(1)
func (ls *linked_list[T]) deleteFront() {
	if ls.head != nil {
		ls.head = ls.head.next
	}
	ls.size--
}

// O(n)
func (ls *linked_list[T]) deleteBack() {
	if ls.head != nil {
		tempNode := ls.head
		for tempNode.next.next != nil {
			tempNode = tempNode.next
		}
		tempNode.next = nil
	}
	ls.size--

}

func (ls *linked_list[T]) getSize() int {
	return ls.size
}

func (ls *linked_list[T]) printList() {
	current := ls.head
	for current != nil {
		fmt.Printf("%v ", current.val)
		current = current.next
	}
	fmt.Println()
}