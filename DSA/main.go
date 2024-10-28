package main






func main(){

var g graph[string]

g.AddVertex("A")
g.AddVertex("B")
g.AddVertex("C")
g.AddVertex("D")
g.AddVertex("E")
g.AddVertex("F")
g.AddVertex("G")



g.AddEdge("A","B")
g.AddEdge("A","C")
g.AddEdge("A","E")
g.AddEdge("B","D")
g.AddEdge("B","F")
g.AddEdge("D","G")
g.AddEdge("C","G")


// g.PrintGraph()
g.DFS("A",make(map[string]bool))




}