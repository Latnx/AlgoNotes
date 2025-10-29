package main

import (
	"fmt"
)

type Edge struct {
	To     int
	Weight int
}
type Graph struct {
	n   int
	adj [][]Edge
}

func NewGraph(n int) *Graph {
	g := &Graph{
		n,
		make([][]Edge, n+1),
	}
	for i := range n + 1 {
		g.adj[i] = make([]Edge, 0)
	}
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.adj[u] = append(g.adj[u], Edge{To: v, Weight: w})
}

func (g *Graph) String() {
	for i := 0; i < g.n; i++ {
		for _, e := range g.adj[i] {
			fmt.Printf("%d:(%d, w=%d) ", i, e.To, e.Weight)
		}
	}
}

func main() {

}
