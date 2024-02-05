package graph

import (
	"strconv"

	bag "github.com/AlbertRossJoh/itualgs_go/fundamentals/bag"
	sharedfunctions "github.com/AlbertRossJoh/itualgs_go/utilities/sharedFunctions"
)

type edgeWeightedGraph struct {
	v   uint
	e   uint
	adj []bag.Bag[Edge]
}

func NewEdgeWeightedGraph(v uint) *edgeWeightedGraph {
	tmp := make([]bag.Bag[Edge], v)
	for i := 0; i < int(v); i++ {
		tmp = append(tmp, *bag.NewBag[Edge]())
	}
	return &edgeWeightedGraph{
		v:   v,
		e:   0,
		adj: tmp,
	}
}

func (graph *edgeWeightedGraph) AddEdge(e Edge) error {
	v := e.Either()
	w, err := e.Other(v)
	if err != nil {
		return err
	}
	graph.validate(v)
	graph.validate(w)
	graph.adj[v].Add(e)
	graph.adj[w].Add(e)
	graph.e++
	return nil
}

func (graph *edgeWeightedGraph) Adj(v uint) *sharedfunctions.Iterator[Edge] {
	graph.validate(v)
	return graph.adj[v].GetIterator()
}

func (graph *edgeWeightedGraph) Degree(v uint) int {
	graph.validate(v)
	return graph.adj[v].Size
}

func (graph *edgeWeightedGraph) Edges() *bag.Bag[Edge] {
	ret := bag.NewBag[Edge]()
	var v uint
	for v = 0; v < graph.v; v++ {
		selfLoops := 0
		curr := graph.adj[v].GetIterator()
		for curr.HasNext() {
			e, _ := curr.Next()
			ev, _ := e.Other(v)
			if ev > v {
				ret.Add(e)
			} else if ev == v {
				if selfLoops%2 == 0 {
					ret.Add(e)
				}
				selfLoops++
			}
		}
	}
	return ret
}

func (graph *edgeWeightedGraph) validate(v uint) {
	if int(v) >= len(graph.adj) {
		panic("vertex " + strconv.Itoa(int(v)) + " is not between 0 and " + strconv.Itoa(len(graph.adj)))
	}
}
