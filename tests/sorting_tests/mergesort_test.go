package tests

import (
	"math/rand"
	"testing"
	"time"

	sort "github.com/AlbertRossJoh/itualgs_go/sorting"
	utils "github.com/AlbertRossJoh/itualgs_go/utilities/sharedFunctions"
)

func TestMergeSort(t *testing.T) {
	rand.New(rand.NewSource(time.Now().Unix()))
	a := rand.Perm(1_000_000)
	sort.MergeSort(&a)

	if !utils.IsSorted(&a) {
		t.Errorf("Mergesort failed")
	}
}
