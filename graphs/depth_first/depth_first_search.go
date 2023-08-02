package depthfirst

import (
	"strconv"

	"github.com/AlbertRossJoh/itualgs_go/graphs/graph"
)

type DFS struct {
	visited []bool
	count   int
}

func NewDFS(g graph.Graph, n int) DFS {
	dfs := DFS{
		visited: make([]bool, n),
		count:   0,
	}
	dfs.DFS(g, n)
	return dfs
}

func (dfs *DFS) DFS(g graph.Graph, v int) {
	dfs.count++
	dfs.visited[v] = true
	iter := g.Adjecent(v)
	for iter.HasNext() {
		val, _ := iter.Next()
		if !dfs.visited[val] {
			dfs.DFS(g, val)
		}
	}
}

func (dfs *DFS) validateVertex(v int) {
	if v < 0 || v >= len(dfs.visited) {
		panic("vertex is not between 0 and " + strconv.Itoa(len(dfs.visited)-1))
	}
}

func (dfs *DFS) IsConnected(v int) bool {
	dfs.validateVertex(v)
	return dfs.visited[v]
}

func (dfs *DFS) Count() int {
	return dfs.count
}
