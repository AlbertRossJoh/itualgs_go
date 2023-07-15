package tests

import (
	"testing"

	fund "github.com/AlbertRossJoh/itualgs_go/fundementals"
)

func TestNewMatrix(t *testing.T) {
	m := fund.NewMatrix(3, 3)
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
	m := fund.CreateFromArray(&a)

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
	m := fund.CreateFromArray(&a)
	k := fund.CreateFromArray(&b)
	m.Transpose()

	if !m.IsEqual(*k) {
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

	m := fund.CreateFromArray(&a)
	k := fund.CreateFromArray(&b)
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

	m := fund.CreateFromArray(&a)

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

	m := fund.CreateFromArray(&a)
	dummy := m
	k := fund.NewIdentityMatrix(3)
	res := m.ComputeInverse().Product(dummy)
	if res.IsEqual(k) {
		panic("Inverse does not work")
	}
}
