package graph

import (
	"strconv"

	"github.com/AlbertRossJoh/itualgs_go/fundamentals/bag"
	util "github.com/AlbertRossJoh/itualgs_go/utilities"
)

type Graph struct {
	v   int
	e   int
	adj []bag.Bag[int]
}

func NewGraph(v int) Graph {
	if v < 0 {
		panic("Number of vertices must be nonnegative")
	}
	return Graph{v, 0, make([]bag.Bag[int], v)}
}

func (g *Graph) Clone() Graph {
	tmpAdj := make([]bag.Bag[int], len(g.adj))
	for i := 0; i < len(g.adj); i++ {
		tmpAdj[i] = g.adj[i].Clone()
	}
	return Graph{g.v, g.e, tmpAdj}
}

func (g *Graph) Vertices() int {
	return g.v
}

func (g *Graph) Edges() int {
	return g.e
}

func validateVertex(g *Graph, v int) {
	if v < 0 || v >= len(g.adj) {
		panic("vertex " + strconv.Itoa(v) + " is not between 0 and " + strconv.Itoa(len(g.adj)))
	}
}

func (g *Graph) AddEdge(v, w int) {
	validateVertex(g, v)
	validateVertex(g, w)
	g.e++
	g.adj[v].Add(w)
	g.adj[w].Add(v)
}

func (g *Graph) Adjecent(v int) util.Iterator[int] {
	validateVertex(g, v)
	return *g.adj[v].GetIterator()
}

func (g *Graph) Degree(v int) int {
	validateVertex(g, v)
	return g.adj[v].Size
}
