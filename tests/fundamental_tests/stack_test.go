package tests

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/AlbertRossJoh/itualgs_go/utilities/sharedFunctions"

	. "github.com/AlbertRossJoh/itualgs_go/fundamentals/stack"
)

func TestStack(t *testing.T) {
	s := NewEmptyStack[int]()
	rand.New(rand.NewSource(time.Now().Unix()))

	testArr := rand.Perm(100)

	expected := GetReversed(&testArr)
	for _, val := range testArr {
		s.Push(val)
	}

	for _, val := range expected {
		res, err := s.Pop()
		if err != nil {
			t.Error("Stack encountered an error: ", err)
		}
		if val != res {
			t.Error("Stack failed, expected ", val, "got ", res)
			t.Error("Values in stack: ", s.Items(), "expected ", expected)
		}
	}

}
