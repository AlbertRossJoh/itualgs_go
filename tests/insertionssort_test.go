package tests

import (
	"reflect"
	"testing"

	sort "github.com/AlbertRossJoh/itualgs_go/sorting"
)

func TestInsertionSort(t *testing.T) {
	a := []int{5, 4, 3, 2, 1}
	sort.InsertionSort(&a)

	if !reflect.DeepEqual(a, []int{1, 2, 3, 4, 5}) {
		t.Errorf("InsertionSort failed")
	}
}
