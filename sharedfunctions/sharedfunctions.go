package sharedfunctions

import (
	"math"

	"golang.org/x/exp/constraints"
)

func IsClose(x float64, y float64) bool {
	return math.Abs(x-y) < 0.000001
}

func IsSorted[T constraints.Ordered](arr *[]T) bool {
	for i := 0; i < len(*arr)-1; i++ {
		if (*arr)[i] > (*arr)[i+1] {
			return false
		}
	}
	return true
}
