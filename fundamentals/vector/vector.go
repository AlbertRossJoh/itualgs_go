package vector

import (
	"math"

	"github.com/AlbertRossJoh/itualgs_go/customerrors"
	utils "github.com/AlbertRossJoh/itualgs_go/utilities/sharedFunctions"
)

type Vector struct {
	dimension int
	Elements  *[]float64
}

func NewVector(dimension int) Vector {
	arr := make([]float64, dimension)
	return Vector{
		dimension: dimension,
		Elements:  &arr,
	}
}

func CreateVectorFromArray(array []float64) Vector {
	if array == nil {
		panic("Nil value provided for array")
	}
	return Vector{
		dimension: len(array),
		Elements:  &array,
	}
}

func (v *Vector) Dimension() int {
	return v.dimension
}

// GetElements Returns a cloned slice of the Elements in a vector
func (v *Vector) GetElements() []float64 {
	clone := make([]float64, v.dimension)
	copy(clone, *v.Elements)
	return clone
}

func (v *Vector) Dot(other *Vector) float64 {
	var result float64
	for i := 0; i < v.dimension; i++ {
		result += (*v.Elements)[i] * (*other.Elements)[i]
	}
	return result
}

func (v *Vector) Add(other Vector) {
	for i := 0; i < v.dimension; i++ {
		(*v.Elements)[i] += (*other.Elements)[i]
	}
}

func (v *Vector) Sub(other Vector) {
	for i := 0; i < v.dimension; i++ {
		(*v.Elements)[i] -= (*other.Elements)[i]
	}
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v *Vector) Normalize() {
	magnitude := v.Magnitude()
	for i := 0; i < v.dimension; i++ {
		(*v.Elements)[i] = (*v.Elements)[i] / magnitude
	}
}

func (v *Vector) DistanceTo(other *Vector) float64 {
	tmp := v
	tmp.Sub(*other)
	return tmp.Magnitude()
}

func (v *Vector) Cartesian(i int) float64 {
	return (*v.Elements)[i]
}

func (v *Vector) Multiply(scalar float64) {
	for i := 0; i < v.dimension; i++ {
		(*v.Elements)[i] = (*v.Elements)[i] * scalar
	}
}

func Multiply(v *Vector, scalar float64) Vector {
	tmp := v.Clone()
	tmp.Multiply(scalar)
	return *tmp
}

// Direction Returns a vector with the same direction as the vector but with a magnitude of 1
func (v *Vector) Direction() (*Vector, error) {
	if v.Magnitude() == 0 {
		return &Vector{}, &customerrors.ErrZeroVector{}
	}
	tmp := v.Clone()
	tmp.Multiply(1 / v.Magnitude())
	return tmp, nil
}

func (v *Vector) Equals(other *Vector) bool {
	if v.dimension != other.dimension {
		return false
	}
	for i := 0; i < v.dimension; i++ {
		if !utils.IsClose((*v.Elements)[i], (*other.Elements)[i]) {
			return false
		}
	}
	return true
}

func (v *Vector) AngleTo(other *Vector) (float64, error) {
	if v.Magnitude() == 0 || other.Magnitude() == 0 {
		return 0, &customerrors.ErrZeroVector{}
	}
	return math.Acos(v.Dot(other) / (v.Magnitude() * other.Magnitude())), nil
}

func (v *Vector) Projection(other Vector) (*Vector, error) {
	if other.Magnitude() == 0 {
		return &Vector{}, &customerrors.ErrZeroVector{}
	}
	tmp := other.Clone()
	tmp.Multiply(other.Dot(v) / other.Dot(&other))
	return tmp, nil
}

func (v *Vector) Cross(other *Vector) (float64, error) {
	if v.dimension > 2 || other.dimension > 2 {
		return 0, &customerrors.ErrVectorCross{}
	}
	return (*v.Elements)[0]*(*other.Elements)[1] - (*v.Elements)[1]*(*other.Elements)[0], nil
}

func (v *Vector) Clone() *Vector {
	clone := make([]float64, v.dimension)
	copy(clone, *v.Elements)
	return &Vector{
		dimension: v.dimension,
		Elements:  &clone,
	}
}
