package graph

import "errors"

type Edge struct {
	v      uint
	w      uint
	Weight int
}

func NewEdge(v, w uint, weight int) *Edge {
	return &Edge{v, w, weight}
}

func (e *Edge) Either() uint {
	return e.v
}

func (e *Edge) Other(vertex uint) (uint, error) {
	if vertex == e.v {
		return e.w, nil
	} else if vertex == e.w {
		return e.v, nil
	} else {
		return 0, errors.New("Other is not valid")
	}
}

func (e *Edge) Cmp(that Edge) int {
	if that.Weight > e.Weight {
		return -1
	} else if that.Weight < e.Weight {
		return 1
	} else {
		return 0
	}
}
