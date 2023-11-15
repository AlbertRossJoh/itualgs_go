package tests

import (
	"math/rand"
	"testing"
	"time"

	sort "github.com/AlbertRossJoh/itualgs_go/sorting"
	utils "github.com/AlbertRossJoh/itualgs_go/utilities/sharedFunctions"
)

func TestQuickSort(t *testing.T) {

	rand.New(rand.NewSource(time.Now().Unix()))
	a := rand.Perm(1_000)
	sort.QuickSort(&a)

	if !utils.IsSorted(&a) {
		t.Errorf("Quicksort failed")
	}
}

func TestQuickSelect(t *testing.T) {

	rand.New(rand.NewSource(time.Now().Unix()))
	a := rand.Perm(1_000)

	if sort.QuickSelect(&a, 25) != 25 {
		t.Errorf("Quickselect failed")
	}
}
