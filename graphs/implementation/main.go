package main

import "fmt"

//     [3] ------- [4] ------- [5]
//      |           |           |
//      |           |           |
//     [1] ------- [2]         [6]
//        \       /
//         \     /
//          \   /
//           [0]

type MyGraph struct {
	adjacentList  map[int][]int
	numberOfNodes int
}

func NewGraph() *MyGraph {
	return &MyGraph{
		adjacentList:  make(map[int][]int),
		numberOfNodes: 0,
	}
}

func (g *MyGraph) AddVertex(node int) {
	g.adjacentList[node] = nil
	g.numberOfNodes++
}

func (g *MyGraph) AddEdge(node1, node2 int) {
	g.adjacentList[node1] = append(g.adjacentList[node1], node2)
}

func main() {
	graph := NewGraph()

	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)
	graph.AddVertex(6)

	graph.AddEdge(0, 1)
	// Since undirected we need to connect both
	graph.AddEdge(1, 0)

	graph.AddEdge(0, 2)
	graph.AddEdge(2, 0)

	graph.AddEdge(1, 3)
	graph.AddEdge(3, 1)

	graph.AddEdge(1, 2)
	graph.AddEdge(2, 1)

	graph.AddEdge(2, 4)
	graph.AddEdge(4, 2)

	graph.AddEdge(3, 4)
	graph.AddEdge(4, 3)

	graph.AddEdge(4, 5)
	graph.AddEdge(5, 4)

	graph.AddEdge(5, 6)
	graph.AddEdge(6, 5)

	fmt.Printf("%+v\n", *graph)
}
