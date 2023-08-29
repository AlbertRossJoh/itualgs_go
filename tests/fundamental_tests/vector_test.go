package tests

import (
	"testing"

	. "github.com/AlbertRossJoh/itualgs_go/fundamentals/vector"
)

func TestNewVector(t *testing.T) {
	arr := []float64{1, 2}
	v := CreateVectorFromArray(arr)

	if (*v.Elements)[0] != 1 || (*v.Elements)[1] != 2 {
		t.Error("Expected 1 and 2, got ", (*v.Elements)[0], (*v.Elements)[1])
	}
}

func TestDotProduct(t *testing.T) {
	arr := []float64{1, 2, 5}
	v := CreateVectorFromArray(arr)
	v2 := CreateVectorFromArray(arr)
	dp := v.Dot(&v2)
	if dp != 30 {
		t.Error("Expected 30, got ", dp)
	}
}

func TestProjection(t *testing.T) {
	a := CreateVectorFromArray(proj_test_vec1)
	b := CreateVectorFromArray(proj_test_vec2)

	res, _ := (&a).Projection(b)

	expected := CreateVectorFromArray(proj_test_res)

	if !res.Equals(&expected) {
		t.Error("Expected ", expected.Elements, " got ", res.Elements)
		t.Fail()
	}
}
