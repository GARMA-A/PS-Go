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



func (g *graph[T]) DFS(start T,visited map[T]bool) {
	visited[start] = true
	fmt.Printf("%v ", start)
	for temp := g.vertices[start].head ;temp != nil; {
		if !visited[temp.val] {
			g.DFS(temp.val, visited)
		}
		temp = temp.next
	}
}


























// // DFS algorithm
// func (g *Graph) DFS(start string, visited map[string]bool) {
// 	// Mark the current node as visited
// 	visited[start] = true
// 	fmt.Printf("%s ", start)

// 	// Recur for all the vertices adjacent to this vertex
// 	for _, neighbor := range g.vertices[start] {
// 		if !visited[neighbor] {
// 			g.DFS(neighbor, visited)
// 		}
// 	}
// }


