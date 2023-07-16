package sharedfunctions

import (
	"math"
	"math/rand"

	"golang.org/x/exp/constraints"
)

func IsClose(x float64, y float64) bool {
	return math.Abs(x-y) < 0.00001
}

func IsSorted[T constraints.Ordered](arr *[]T) bool {
	for i := 0; i < len(*arr)-1; i++ {
		if (*arr)[i] > (*arr)[i+1] {
			return false
		}
	}
	return true
}

func CompareArrays[T comparable](arr1 *[]T, arr2 *[]T) bool {
	if len(*arr1) != len(*arr2) {
		return false
	}

	for i := 0; i < len(*arr1); i++ {
		if (*arr1)[i] != (*arr2)[i] {
			return false
		}
	}
	return true
}

func Exchange[T any](arr *[]T, i int, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

func Shuffle[T any](arr *[]T) {
	for i := 0; i < len(*arr); i++ {
		j := rand.Intn(i + 1)
		Exchange(arr, i, j)
	}
}
