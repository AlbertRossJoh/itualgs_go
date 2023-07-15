package fundementals

import (
	"math"

	"github.com/AlbertRossJoh/itualgs_go/customerrors"
)

type Vector struct {
	dimension int
	elements  []float64
}

func NewVector(dimension int) Vector {
	return Vector{
		dimension: dimension,
		elements:  make([]float64, dimension),
	}
}

func (v Vector) Dimension() int {
	return v.dimension
}

func (v Vector) Elements() []float64 {
	return v.elements
}

func (v Vector) Dot(other Vector) float64 {
	var result float64
	for i := 0; i < v.dimension; i++ {
		result += v.elements[i] * other.elements[i]
	}
	return result
}

func (v Vector) Add(other Vector) Vector {
	var result Vector
	result.dimension = v.dimension
	result.elements = make([]float64, v.dimension)
	for i := 0; i < v.dimension; i++ {
		result.elements[i] = v.elements[i] + other.elements[i]
	}
	return result
}

func (v Vector) Sub(other Vector) Vector {
	var result Vector
	result.dimension = v.dimension
	result.elements = make([]float64, v.dimension)
	for i := 0; i < v.dimension; i++ {
		result.elements[i] = v.elements[i] - other.elements[i]
	}
	return result
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vector) Normalize() Vector {
	var result Vector
	result.dimension = v.dimension
	result.elements = make([]float64, v.dimension)
	magnitude := v.Magnitude()
	for i := 0; i < v.dimension; i++ {
		result.elements[i] = v.elements[i] / magnitude
	}
	return result
}

func (v Vector) DistanceTo(other Vector) float64 {
	return v.Sub(other).Magnitude()
}

func (v Vector) Cartesian(i int) float64 {
	return v.elements[i]
}

func (v Vector) Multiply(scalar float64) Vector {
	var result Vector
	result.dimension = v.dimension
	result.elements = make([]float64, v.dimension)
	for i := 0; i < v.dimension; i++ {
		result.elements[i] = v.elements[i] * scalar
	}
	return result
}

func (v Vector) Direction() (Vector, error) {
	if v.Magnitude() == 0 {
		return Vector{}, &customerrors.ErrZeroVector{}
	}
	return v.Multiply(1 / v.Magnitude()), nil
}

func (v Vector) Equals(other Vector) bool {
	if v.dimension != other.dimension {
		return false
	}
	for i := 0; i < v.dimension; i++ {
		if v.elements[i] != other.elements[i] {
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

func (v Vector) Projection(other Vector) (Vector, error) {
	if other.Magnitude() == 0 {
		return Vector{}, &customerrors.ErrZeroVector{}
	}
	return other.Multiply(v.Dot(other) / math.Pow(other.Magnitude(), 2)), nil
}

func (v Vector) Cross(other Vector) (float64, error) {
	if v.dimension > 2 || other.dimension > 2 {
		return 0, &customerrors.ErrVectorCross{}
	}
	return v.elements[0]*other.elements[1] - v.elements[1]*other.elements[0], nil
}
