package tests

import (
	"testing"

	fund "github.com/AlbertRossJoh/itualgs_go/fundamentals"
)

func TestNewVector(t *testing.T) {
	arr := []float64{1, 2}
	v := fund.CreateVectorFromArray(arr)

	if (*v.Elements())[0] != 1 || (*v.Elements())[1] != 2 {
		t.Error("Expected 1 and 2, got ", (*v.Elements())[0], (*v.Elements())[1])
	}
}

func TestDotProduct(t *testing.T) {
	arr := []float64{1, 2, 5}
	v := fund.CreateVectorFromArray(arr)
	v2 := fund.CreateVectorFromArray(arr)
	dp := v.Dot(v2)
	if dp != 30 {
		t.Error("Expected 30, got ", dp)
	}
}
