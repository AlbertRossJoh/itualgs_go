package tests

import (
	"reflect"
	"testing"

	sort "github.com/AlbertRossJoh/itualgs_go/sorting"
)

func TestMergeSort(t *testing.T) {
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.MergeSort(&a)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if !reflect.DeepEqual(a, expected) {
		t.Errorf("InsertionSort failed expected %v got %v", expected, a)
	}
}
