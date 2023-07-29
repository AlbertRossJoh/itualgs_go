package tests

import (
	"math/rand"
	"testing"
	"time"

	utils "github.com/AlbertRossJoh/itualgs_go/sharedfunctions"
	sort "github.com/AlbertRossJoh/itualgs_go/sorting"
)

func TestQuickSort(t *testing.T) {

	rand.Seed(time.Now().Unix())
	a := rand.Perm(1_000)
	sort.QuickSort(&a)

	if !utils.IsSorted(&a) {
		t.Errorf("Quicksort failed")
	}
}

func TestQuickSelect(t *testing.T) {

	rand.Seed(time.Now().Unix())
	a := rand.Perm(1_000)

	if sort.QuickSelect(&a, 25) != 25 {
		t.Errorf("Quickselect failed")
	}
}
