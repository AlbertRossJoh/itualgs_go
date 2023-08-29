package tests

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/AlbertRossJoh/itualgs_go/utilities"
)

func TestIsClose(t *testing.T) {
	a := 0.0000001
	b := 0.0000002

	if !IsClose(a, b) {
		t.Error("Expected", a, "and", b, "to be close")
	}
}

func TestReverse(t *testing.T) {
	rand.New(rand.NewSource(time.Now().Unix()))

	testArr := rand.Perm(100)

	clone := make([]int, len(testArr))
	copy(clone, testArr)
	Reverse(&clone)
	for i := 0; i < len(clone); i++ {
		if clone[i] != testArr[len(testArr)-1-i] {
			t.Error("Expected", testArr, "to be reversed to", clone)
		}
	}
}

func TestGetReversed(t *testing.T) {
	rand.New(rand.NewSource(time.Now().Unix()))

	testArr := rand.Perm(100)

	clone := GetReversed(&testArr)

	for i := 0; i < len(clone); i++ {
		if clone[i] != testArr[len(testArr)-1-i] {
			t.Error("Expected", testArr, "to be reversed to", clone)
		}
	}
}

func TestZip(t *testing.T) {
	rand.New(rand.NewSource(time.Now().Unix()))

	a := rand.Perm(10)
	b := rand.Perm(10)

	clone, err := Zip(&a, &b)

	if err != nil {
		t.Error(err)
	}

	for i := 0; i < len(clone); i += 2 {
		if clone[i] != a[i/2] || clone[i+1] != b[i/2] {
			t.Error("Expected", a, "and", b, "to be zipped to", clone)
			break
		}
	}
}

func TestShuffle(t *testing.T) {
	rand.New(rand.NewSource(time.Now().Unix()))

	testArr := rand.Perm(100)
	clone := make([]int, len(testArr))
	copy(clone, testArr)
	Shuffle(&clone)

	if CompareArrays[int](&testArr, &clone) {
		t.Error("Expected", testArr, "to be shuffled", clone)
		t.Fail()
	}
}

func TestIsSorted(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	if !IsSorted(&a) {
		t.Error("Expected", a, "to be sorted")
	}
}
