package main

import (
	"errors"
	"fmt"
)

type graph[T comparable] struct {
	vertices map[T]*linked_list[T]
}

func (g *graph[T]) AddVertex(vertex T) {
	if g.vertices == nil {
		g.vertices = make(map[T]*linked_list[T])
	}
	g.vertices[vertex] = &linked_list[T]{}
}

func (g *graph[T]) AddEdge(vertex1, vertex2 T) error {
	if _, ok := g.vertices[vertex1]; !ok {
		return errors.New("vertex1 does not exist")
	}
	if _, ok := g.vertices[vertex2]; !ok {
		return errors.New("vertex2 does not exist")
	}
	g.vertices[vertex1].addBack(vertex2)
	// g.vertices[vertex2].addBack(vertex1)
	return nil
}

func (g *graph[T]) PrintGraph() {
	for vertex, edges := range g.vertices {
		fmt.Printf("%v -> ", vertex)
		edges.printList()
	}
	fmt.Println()
}

func (g *graph[T]) DFS(start T, visited map[T]bool) {
	visited[start] = true
	fmt.Printf("%v ", start)
	for temp := g.vertices[start].head; temp != nil; {
		if !visited[temp.val] {
			g.DFS(temp.val, visited)
		}
		temp = temp.next
	}
}

func (g *graph[T]) allVerticesOnReverseOrder(vertex T) []T {
	if linkedList, ok := g.vertices[vertex]; ok {
		size := linkedList.getSize()
		sl := make([]T, size)

		node := linkedList.head
		for i := size - 1; node != nil; i-- {
			sl[i] = node.val
			node = node.next
		}

		return sl
	}
	return []T{}
}

func (g *graph[T]) IterativeDFS(start T) {

	var visited = map[T]bool{}
	var st = stack[T]{}
	st.push(start)
	for st.size != 0 {
		topNode := st.pop().val
		if !visited[topNode] {
			fmt.Print(topNode, " ")
			visited[topNode] = true
		}
		neigbors := g.allVerticesOnReverseOrder(topNode)
		for _, vertex := range neigbors {
			if !visited[vertex] {
				st.push(vertex)
			}
		}
	}

}

func (g *graph[T]) IterativeBFS(start T) {
	visited := map[T]bool{}
	q := queue[T]{}
	q.inqueue(start)
	for q.size != 0 {
		front := q.front.val
		q.dequeue()
		fmt.Print(front, " ")
		neighbors := g.allVerticesOnReverseOrder(front)
		for _, edge := range neighbors {
			if !visited[edge] {
				q.inqueue(edge)
				visited[edge] = true
			}
		}

	}
}

func (g *graph[T]) BFS(q *queue[T], visited map[T]bool) {

	if q.size == 0 {
		return
	}
	front := q.front.val
	q.dequeue()
	fmt.Print(front, " ")

	edges := g.allVerticesOnReverseOrder(front)
	for _, edge := range edges {
		if !visited[edge] {
			q.inqueue(edge)
			visited[edge] = true
		}
	}
	g.BFS(q, visited)

}
