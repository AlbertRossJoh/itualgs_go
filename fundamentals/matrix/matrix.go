package fundementals

import (
	"math"

	. "github.com/AlbertRossJoh/itualgs_go/fundamentals/vector"
	util "github.com/AlbertRossJoh/itualgs_go/utilities"
)

type Matrix struct {
	Rows int
	Cols int
	Data *[][]float64
}

func CreateMatrixFromArray(arr *[][]float64) *Matrix {
	if arr == nil {
		panic("Array is nil")
	}
	return &Matrix{
		Rows: len(*arr),
		Cols: len((*arr)[0]),
		Data: arr,
	}
}

func CreateMatrixFromArrayOfVectors(arr []Vector) *Matrix {
	if arr == nil {
		panic("Array is nil")
	}
	acc := make([][]float64, 0, len(arr))
	for _, v := range arr {
		acc = append(acc, *v.Elements)
	}
	return &Matrix{
		Rows: len(arr),
		Cols: arr[0].Dimension(),
		Data: &acc,
	}
}

func NewMatrix(rows, cols int) Matrix {
	m := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]float64, cols)
	}
	return Matrix{
		Rows: rows,
		Cols: cols,
		Data: &m,
	}
}

func NewIdentityMatrix(n int) Matrix {
	m := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		(*m.Data)[i][i] = 1
	}
	return m
}

func (m *Matrix) AugmentRight(v Vector) {
	if v.Dimension() != m.Rows {
		panic("Dimension mismatch")
	}
	tmp := NewMatrix(m.Rows, m.Cols+1)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			(*tmp.Data)[i][j] = (*m.Data)[i][j]
		}
		(*tmp.Data)[i][m.Cols] = (*v.Elements)[i]
	}
	m.Rows, m.Cols, m.Data = tmp.Rows, tmp.Cols, tmp.Data
}

func (m *Matrix) MatrixVectorProduct(v Vector) Vector {
	tmp := NewVector(m.Rows)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			(*tmp.Elements)[i] += (*m.Data)[i][j] * (*v.Elements)[j]
		}
	}
	return tmp
}

func (m *Matrix) Transpose() Matrix {
	tmp := NewMatrix(m.Cols, m.Rows)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			(*tmp.Data)[j][i] = (*m.Data)[i][j]
		}
	}
	return tmp
}

func (m *Matrix) Product(n *Matrix) *Matrix {
	tmp := NewMatrix(m.Rows, n.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < n.Cols; j++ {
			for k := 0; k < n.Rows; k++ {
				(*tmp.Data)[i][j] += (*m.Data)[i][k] * (*n.Data)[k][j]
			}
		}
	}
	return &tmp
}

func (m *Matrix) IsSquare() bool {
	return m.Rows == m.Cols
}

func (m *Matrix) Power(power int) {
	acc := m
	for i := 0; i < power-1; i++ {
		acc = m.Product(acc)
	}
	m.Data = acc.Data
}

func (m *Matrix) RowReplacement(i int, j int, num float64) {
	if i != j {
		for k := 0; k < m.Cols; k++ {
			(*m.Data)[i][k] = (*m.Data)[j][k]*num + (*m.Data)[i][k]
		}
	}
}

func (m *Matrix) RowInterchange(row1 int, row2 int) {
	acc := make([]float64, m.Cols)
	for k := 0; k < m.Cols; k++ {
		acc[k] = (*m.Data)[row1][k]
		(*m.Data)[row1][k] = (*m.Data)[row2][k]
		(*m.Data)[row2][k] = acc[k]
	}
}

func (m *Matrix) RowScaling(row int, num float64) {
	for k := 0; k < m.Cols; k++ {
		(*m.Data)[row][k] *= num
	}
}

func (m *Matrix) ForwardReduction() {
	row := 0
	for i := 0; i < m.Cols; i++ {
		for j := row; j < m.Rows; j++ {
			if !util.IsClose((*m.Data)[row][i], 0) && m.isZeroColumn(i) {
				m.RowInterchange(row, j)
				for k := j + 1; k < m.Rows; k++ {
					m.RowReplacement(k, j, -(*m.Data)[k][i]/(*m.Data)[row][i])
				}
				row++
				break
			}
		}
	}
}

func (m *Matrix) BackwardReduction() {
	for i := m.Rows - 1; i >= 0; i-- {
		for j := 0; j < m.Cols; j++ {
			if !util.IsClose((*m.Data)[i][j], 0) {
				m.RowScaling(i, 1/(*m.Data)[i][j])
				for k := i - 1; k >= 0; k-- {
					m.RowReplacement(k, i, -(*m.Data)[k][j]/(*m.Data)[i][j])
				}
				break
			}
		}
	}
}

func (m Matrix) isZeroColumn(col int) bool {
	for i := 0; i < m.Rows; i++ {
		if (*m.Data)[i][col] != 0 {
			return true
		}
	}
	return false
}

func (m *Matrix) GaussElimination(v Vector) Vector {
	tmp := m
	tmp.AugmentRight(v)
	tmp.ForwardReduction()
	tmp.BackwardReduction()
	return tmp.ExtractColumn(m.Cols - 1)
}

func (m Matrix) ExtractColumn(col int) Vector {
	if col >= m.Cols {
		panic("Column out of bounds")
	}
	acc := NewVector(m.Rows)
	for i := 0; i < m.Rows; i++ {
		(*acc.Elements)[i] = (*m.Data)[i][col]
	}
	return acc
}

func (m *Matrix) ComputeInverse() *Matrix {
	if !m.IsSquare() {
		panic("Matrix is not square")
	}
	identity := NewIdentityMatrix(m.Rows)
	dummy := clone(m)
	row := 0
	for i := 0; i < m.Cols; i++ {
		for j := row; j < m.Rows; j++ {
			if !util.IsClose((*dummy.Data)[j][i], 0) {
				dummy.RowInterchange(row, j)
				identity.RowInterchange(row, j)
				for k := j + 1; k < m.Rows; k++ {
					scale := -(*dummy.Data)[k][i] / (*dummy.Data)[row][i]
					dummy.RowReplacement(k, j, scale)
					identity.RowReplacement(k, j, scale)
				}
				row++
				break
			}
		}
	}
	for i := m.Rows - 1; i >= 0; i-- {
		for j := 0; j < m.Cols; j++ {
			if !util.IsClose((*dummy.Data)[i][j], 0) {
				scale := 1 / (*dummy.Data)[i][j]
				dummy.RowScaling(i, scale)
				identity.RowScaling(i, scale)
				for k := i; k >= 0; k-- {
					scale = -(*dummy.Data)[k][j] / (*dummy.Data)[i][j]
					dummy.RowReplacement(k, i, scale)
					identity.RowReplacement(k, i, scale)
				}
				break
			}
		}
	}
	return &identity
}

func (m Matrix) IsEqual(n Matrix) bool {
	if m.Rows != n.Rows || m.Cols != n.Cols {
		return false
	}
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if !util.IsClose((*m.Data)[i][j], (*n.Data)[i][j]) {
				return false
			}
		}
	}
	return true
}

func subMatrix(m *Matrix, row int, col int) *Matrix {
	if !m.IsSquare() {
		panic("Matrix is not square")
	}
	tmp := NewMatrix(m.Rows-1, m.Cols-1)
	acc := []float64{}
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if i != row && j != col {
				acc = append(acc, (*m.Data)[i][j])
			}
		}
	}
	num := 0
	for i := 0; i < tmp.Rows; i++ {
		for j := 0; j < tmp.Cols; j++ {
			(*tmp.Data)[i][j] = acc[num]
			num++
		}
	}
	return &tmp
}

func (m *Matrix) Determinant() float64 {
	var ans float64
	if m.Rows == 2 {
		return (*m.Data)[0][0]*(*m.Data)[1][1] - (*m.Data)[0][1]*(*m.Data)[1][0]
	}
	for i := 0; i < m.Rows; i++ {
		sub := subMatrix(m, 0, i)
		det := sub.Determinant()
		ans += math.Pow(-1, float64(i+2)) * (*m.Data)[0][i] * det
	}
	return ans
}

func clone(m *Matrix) *Matrix {
	tmp := NewMatrix(m.Rows, m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			(*tmp.Data)[i][j] = (*m.Data)[i][j]
		}
	}
	return &tmp
}

func (m *Matrix) OrthogonalBasis() Matrix {
	B := NewMatrix(m.Rows, m.Cols)
	acc := make([]Vector, 0, m.Cols)
	first := NewVector(m.Rows)
	for i := 0; i < m.Rows; i++ {
		(*first.Elements)[i] = (*m.Data)[i][0]
		(*B.Data)[i][0] = (*m.Data)[i][0]
	}
	acc = append(acc, first)
	for i := 1; i < m.Cols; i++ {
		current := NewVector(m.Rows)
		for j := 0; j < m.Rows; j++ {
			(*current.Elements)[j] = (*m.Data)[j][i]
		}
		subVector := NewVector(m.Rows)
		for _, val := range acc {
			if !util.IsClose((&current).Dot(val), 0) {
				val.Multiply(-(val.Dot(current)) / val.Dot(val))
				subVector.Add(val)
			}
		}
		current.Add(subVector)
		acc = append(acc, current)
		for j := 0; j < m.Rows; j++ {
			(*B.Data)[j][i] = (*current.Elements)[j]
		}
	}
	return B
}

func (m *Matrix) GramSchmidt() (Matrix, Matrix) {
	Q := m.OrthogonalBasis()
	Q = Q.Transpose()
	for i := 0; i < Q.Rows; i++ {
		var acc float64
		for j := 0; j < Q.Cols; j++ {
			acc += math.Pow((*Q.Data)[i][j], 2)
		}
		Q.RowScaling(i, 1/math.Sqrt(acc))
	}
	return Q.Transpose(), *Q.Product(m)
}

func (m *Matrix) NonNormalGS() Matrix {
	tmp := *m
	for i := 0; i < tmp.Rows; i++ {
		var acc float64
		for j := 0; j < tmp.Cols; j++ {
			acc += math.Pow((*tmp.Data)[i][j], 2)
		}
		tmp.RowScaling(i, 1/math.Sqrt(acc))
	}
	return tmp
}

func vecArrGS(m *[]Vector) []*Vector {
	var ret []*Vector
	for _, valI := range *m {
		row := valI.Clone()
		for _, vec := range ret {
			proj, err := vec.Projection(valI)
			if err != nil {
				panic(err)
			}
			row.Sub(*proj)
		}
		if row.Magnitude() != 0 {
			ret = append(ret, row)
		}
	}
	return ret
}

// Something is wrong with this function do not use
func (m *Matrix) LLL() Matrix {
	base := make([]Vector, 0, m.Rows)
	for _, row := range *m.Data {
		clone := make([]float64, len(row))
		copy(clone, row)
		base = append(base, CreateVectorFromArray(clone))
	}
	res := vecArrGS(&base)
	k := 1

	var mu = func(i int, j int) float64 {
		valI := base[i]
		valJ := base[j]
		return (valI.Dot(valJ)) / (valI.Dot(valI))
	}

	for k < len(res) {
		for j := k - 1; j > -1; j-- {
			currMU := mu(k, j)
			if math.Abs(currMU) > 0.5 {
				base[j].Multiply(currMU)
				base[k].Sub(base[j])
				res = vecArrGS(&base)
			}
		}
		if base[k].Dot(base[k]) > (0.75-math.Pow(mu(k, k-1), 2))*base[k-1].Dot(base[k-1]) {
			k++
		} else {
			base[k], base[k-1] = base[k-1], base[k]
			res = vecArrGS(&base)
			k = int(math.Max(float64(k-1), 2))
		}
	}

	return *CreateMatrixFromArrayOfVectors(base)
}
