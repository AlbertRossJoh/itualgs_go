package tests

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/AlbertRossJoh/itualgs_go/fundamentals/queue"
)

func TestQueue(t *testing.T) {
	s := NewEmptyQueue[int]()
	rand.New(rand.NewSource(time.Now().Unix()))
	testArr := rand.Perm(100)

	for _, val := range testArr {
		s.Enqueue(val)
	}
	for _, val := range testArr {
		res, err := s.Dequeue()
		if err != nil {
			t.Error("Stack encountered an error: ", err)
		}
		if val != res {
			t.Error("Stack failed, expected ", testArr, "got ", s.Items(), "")
		}
	}

}
