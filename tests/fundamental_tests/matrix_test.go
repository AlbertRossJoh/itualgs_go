package tests

import (
	"testing"

	. "github.com/AlbertRossJoh/itualgs_go/fundamentals/matrix"
	. "github.com/AlbertRossJoh/itualgs_go/fundamentals/vector"
)

func TestNewMatrix(t *testing.T) {
	m := NewMatrix(3, 3)
	(*m.Data)[0][0] = 1
	(*m.Data)[0][1] = 1
	if !m.IsSquare() {
		t.Fatal()
		panic("Matrix is not square")
	}

	if (*m.Data)[0][0] != 1 {
		t.Fatal()
		panic("Matrix is not initialized correctly")
	}
}

func TestNewMatrix2(t *testing.T) {

	a := [][]float64{
		{1, 2, -5},
		{5, -2, -13},
		{-3, 3, 15},
		{1, -1, 7},
	}
	m := CreateMatrixFromArray(&a)

	(*m.Data)[0][0] = 1
	(*m.Data)[0][1] = 5
	if m.IsSquare() {
		t.Fatal()
		panic("Matrix is square")
	}

	if (*m.Data)[0][0] != 1 {
		t.Fatal()
		panic("Matrix is not initialized correctly")
	}
}

func TestTranspose(t *testing.T) {

	a := [][]float64{
		{1, 2, -5},
		{5, -2, -13},
		{-3, 3, 15},
		{1, -1, 7},
	}
	b := [][]float64{
		{1, 5, -3, 1},
		{2, -2, 3, -1},
		{-5, -13, 15, 7},
	}
	m := CreateMatrixFromArray(&a)
	k := CreateMatrixFromArray(&b)
	res := m.Transpose()

	if !res.IsEqual(*k) {
		panic("Transpose does not work")
	}
}

func TestPower(t *testing.T) {

	a := [][]float64{
		{1, 5, -3},
		{2, -2, 3},
		{-5, -13, 15},
	}
	b := [][]float64{
		{26, 34, -33},
		{-17, -25, 33},
		{-106, -194, 201},
	}

	m := CreateMatrixFromArray(&a)
	k := CreateMatrixFromArray(&b)
	m.Power(2)

	if !m.IsEqual(*k) {
		panic("Transpose does not work")
	}
}

func TestDeterminant(t *testing.T) {

	a := [][]float64{
		{1, 5, -3},
		{2, -2, 3},
		{-5, -13, 15},
	}

	m := CreateMatrixFromArray(&a)

	if m.Determinant() != -108 {
		panic("Transpose does not work")
	}
}

func TestInverse(t *testing.T) {

	a := [][]float64{
		{1, 5, -3},
		{2, -2, 3},
		{-5, -13, 15},
	}

	m := CreateMatrixFromArray(&a)
	dummy := *m
	k := NewIdentityMatrix(3)
	res := (m.ComputeInverse())
	res = res.Product(&dummy)
	if !res.IsEqual(k) {
		t.Errorf("Inverse does not work, expected %v got %v", k.Data, res.Data)
	}
}

func TestProduct(t *testing.T) {
	a := CreateMatrixFromArray(&prod_testa)
	b := CreateMatrixFromArray(&prod_testb)
	expected := CreateMatrixFromArray(&prod_test_res)

	res := a.Product(b)

	if !res.IsEqual(*expected) {
		panic("Product does not work")
	}
}

func TestMatrixVectorProduct(t *testing.T) {
	a := CreateMatrixFromArray(&mat_vec_prod_test_mat)
	b := CreateVectorFromArray(mat_vec_prod_test_vec)
	expected := CreateVectorFromArray(mat_vec_prod_test_res)
	res := a.MatrixVectorProduct(b)
	if !res.Equals(expected) {
		t.Error("Matrix vector product does not work, expected ", expected.GetElements(), " got ", res.GetElements())
	}
}

func TestMatrixRowReplacement(t *testing.T) {
	a := CreateMatrixFromArray(&mat_row_replace_test_mat)
	res := CreateMatrixFromArray(&mat_row_replace_test_res)
	a.RowReplacement(
		int(mat_row_replace_test_rows[0]),
		int(mat_row_replace_test_rows[1]),
		mat_row_replace_test_val)

	if !a.IsEqual(*res) {
		t.Error("Matrix row replacement does not work, expected ", res.Data, " got ", a.Data)
	}
}

func TestMatrixRowInterchange(t *testing.T) {
	a := CreateMatrixFromArray(&mat_row_interchange_test_mat)
	res := CreateMatrixFromArray(&mat_row_interchange_test_res)
	a.RowInterchange(int(mat_row_interchange_test_rows[0]), int(mat_row_interchange_test_rows[1]))

	if !a.IsEqual(*res) {
		t.Error("Matrix row interchange does not work, expected ", res.Data, " got ", a.Data)
	}
}

func TestMatrixAugmentRight(t *testing.T) {
	a := CreateMatrixFromArray(&aug_right_test_mat)
	b := CreateVectorFromArray(aug_right_test_vec)

	expected := CreateMatrixFromArray(&aug_right_test_res)

	a.AugmentRight(b)

	var collected []float64

	for i := 0; i < a.Rows; i++ {
		collected = append(collected, (*a.Data)[i][a.Cols-1])
	}

	if !a.IsEqual(*expected) {
		t.Error("Matrix augment right does not work, expected ", b.GetElements(), " got ", collected)
	}
}

func TestMatrixForwardReduction(t *testing.T) {
	a := CreateMatrixFromArray(&forward_reduction_test_mat)

	expected := CreateMatrixFromArray(&forward_reduction_test_res)

	a.ForwardReduction()

	if !a.IsEqual(*expected) {
		t.Error("Matrix forward reduction does not work, expected ", expected.Data, " got ", a.Data)
	}
}

func TestMatrixGaussianElimination(t *testing.T) {
	a := CreateMatrixFromArray(&gauss_test_mat)
	b := CreateVectorFromArray(gauss_test_vec)
	expected := CreateVectorFromArray(gauss_test_res)

	res := a.GaussElimination(b)

	if !res.Equals(expected) {
		t.Error("Gaussian elimination does not work, expected ", expected.GetElements(), " got ", res.GetElements())
	}
}

func TestGramSchmidt(t *testing.T) {
	a := CreateMatrixFromArray(&gram_schmidt_test_mat)
	expected1 := CreateMatrixFromArray(&gram_schmidt_test_res_1)
	expected2 := CreateMatrixFromArray(&gram_schmidt_test_res_2)
	res1, res2 := a.GramSchmidt()

	if !res1.IsEqual(*expected1) {
		t.Error("Gram schmidt did not work, expected this ortonormal matrix ", expected1.Data, " got ", res1.Data)
	}

	if !res2.IsEqual(*expected2) {
		t.Error("Gram schmidt did not work, expected this upper matrix ", expected2.Data, " got ", res2.Data)
	}
}

func TestLLL(t *testing.T) {
	a := CreateMatrixFromArray(&lll_test_mat)

	b := CreateMatrixFromArray(&lll_test2_mat)

	expected := CreateMatrixFromArray(&lll_test_res)

	expected2 := CreateMatrixFromArray(&lll_test2_res)

	res := a.LLL()

	res2 := b.LLL()

	if !a.IsEqual(res) {
		t.Error("LLL did not work, expected ", expected, " got ", res)
	}
	if !b.IsEqual(res2) {
		t.Error("LLL did not work, expected ", expected2, " got ", res2)
	}
}
