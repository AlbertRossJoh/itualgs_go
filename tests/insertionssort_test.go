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

func TestInsertionRangeSort(t *testing.T) {
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.InsertionRangeSort(&a, 5, 10)

	if !reflect.DeepEqual(a, []int{10, 9, 8, 7, 6, 1, 2, 3, 4, 5}) {
		t.Errorf("InsertionSort failed expected %v got %v", []int{10, 9, 8, 7, 6, 1, 2, 3, 4, 5}, a)
	}
}

func TestInsertionRangeSort2(t *testing.T) {
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.InsertionRangeSort(&a, 0, len(a))

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if !reflect.DeepEqual(a, expected) {
		t.Errorf("InsertionSort failed expected %v got %v", expected, a)
	}
}
