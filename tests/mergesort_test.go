package tests

import (
	"math/rand"
	"testing"
	"time"

	utils "github.com/AlbertRossJoh/itualgs_go/sharedfunctions"
	sort "github.com/AlbertRossJoh/itualgs_go/sorting"
)

func TestMergeSort(t *testing.T) {

	rand.Seed(time.Now().Unix())
	a := rand.Perm(100_000_000)
	sort.MergeSort(&a)

	if !utils.IsSorted(&a) {
		t.Errorf("Mergesort failed")
	}
}
