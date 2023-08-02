package depthfirst

import (
	"strconv"

	. "github.com/AlbertRossJoh/itualgs_go/customerrors"
	. "github.com/AlbertRossJoh/itualgs_go/fundamentals/stack"
	. "github.com/AlbertRossJoh/itualgs_go/graphs/graph"
	. "github.com/AlbertRossJoh/itualgs_go/utilities"
)

type DFP struct {
	visited []bool
	edgeTo  []int
	s       int
}

func NewDFP(g Graph, s int) *DFP {
	dfp := &DFP{
		visited: make([]bool, g.Vertices()),
		edgeTo:  make([]int, g.Vertices()),
		s:       s,
	}
	dfp.dfs(g, s)
	return dfp
}

func (dfp *DFP) dfs(g Graph, v int) {
	dfp.visited[v] = true
	iter := g.Adjecent(v)
	for iter.HasNext() {
		val, _ := iter.Next()
		if !dfp.visited[val] {
			dfp.edgeTo[val] = v
			dfp.dfs(g, val)
		}
	}
}

func (dfp *DFP) HasPathTo(v int) bool {
	return dfp.visited[v]
}

func (dfp *DFP) PathTo(v int) (*Iterator[int], error) {
	dfp.validateVertex(v)
	stack := NewEmptyStack[int]()
	if !dfp.HasPathTo(v) {
		return stack.GetIterator(), &ErrNonExistantPath{}
	}
	for v != dfp.s {
		stack.Push(v)
		v = dfp.edgeTo[v]
	}
	stack.Push(dfp.s)
	return stack.GetIterator(), nil
}

func (dfp *DFP) validateVertex(v int) {
	if v < 0 || v >= len(dfp.visited) {
		panic("vertex is not between 0 and " + strconv.Itoa(len(dfp.visited)-1))
	}
}
