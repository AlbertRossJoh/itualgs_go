package fundementals

import (
	"math"

	"github.com/AlbertRossJoh/itualgs_go/customerrors"
)

type Vector struct {
	dimension int
	elements  *[]float64
}

func NewVector(dimension int) Vector {
	arr := make([]float64, dimension)
	return Vector{
		dimension: dimension,
		elements:  &arr,
	}
}

func NewVectorFromArray(array *[]float64) Vector {
	if array == nil {
		panic("Nil value provided for array")
	}
	return Vector{
		dimension: len(*array),
		elements:  array,
	}
}

func (v *Vector) Dimension() int {
	return v.dimension
}

func (v *Vector) Elements() *[]float64 {
	return v.elements
}

func (v *Vector) Dot(other Vector) float64 {
	var result float64
	for i := 0; i < v.dimension; i++ {
		result += (*v.elements)[i] * (*other.elements)[i]
	}
	return result
}

func (v *Vector) Add(other Vector) {
	for i := 0; i < v.dimension; i++ {
		(*v.elements)[i] += (*other.elements)[i]
	}
}

func (v *Vector) Sub(other Vector) {
	for i := 0; i < v.dimension; i++ {
		(*v.elements)[i] -= (*other.elements)[i]
	}
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.Dot(*v))
}

func (v *Vector) Normalize() {
	magnitude := v.Magnitude()
	for i := 0; i < v.dimension; i++ {
		(*v.elements)[i] = (*v.elements)[i] / magnitude
	}
}

func (v *Vector) DistanceTo(other *Vector) float64 {
	tmp := v
	tmp.Sub(*other)
	return tmp.Magnitude()
}

func (v *Vector) Cartesian(i int) float64 {
	return (*v.elements)[i]
}

func (v Vector) Multiply(scalar float64) {
	for i := 0; i < v.dimension; i++ {
		(*v.elements)[i] = (*v.elements)[i] * scalar
	}
}

func (v Vector) Direction() (*Vector, error) {
	if v.Magnitude() == 0 {
		return &Vector{}, &customerrors.ErrZeroVector{}
	}
	tmp := v
	tmp.Multiply(1 / v.Magnitude())
	return &tmp, nil
}

func (v Vector) Equals(other Vector) bool {
	if v.dimension != other.dimension {
		return false
	}
	for i := 0; i < v.dimension; i++ {
		if (*v.elements)[i] != (*other.elements)[i] {
			return false
		}
	}
	return true
}

func (v Vector) AngleTo(other Vector) (float64, error) {
	if v.Magnitude() == 0 || other.Magnitude() == 0 {
		return 0, &customerrors.ErrZeroVector{}
	}
	return math.Acos(v.Dot(other) / (v.Magnitude() * other.Magnitude())), nil
}

func (v Vector) Projection(other Vector) (*Vector, error) {
	if other.Magnitude() == 0 {
		return &Vector{}, &customerrors.ErrZeroVector{}
	}
	tmpOther := other
	tmpOther.Multiply(v.Dot(other) / math.Pow(other.Magnitude(), 2))
	return &tmpOther, nil
}

func (v Vector) Cross(other Vector) (float64, error) {
	if v.dimension > 2 || other.dimension > 2 {
		return 0, &customerrors.ErrVectorCross{}
	}
	return (*v.elements)[0]*(*other.elements)[1] - (*v.elements)[1]*(*other.elements)[0], nil
}
