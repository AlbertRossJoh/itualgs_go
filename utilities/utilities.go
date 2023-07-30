package sharedfunctions

import (
	"errors"
	"math"
	"math/rand"
	"time"

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
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(*arr); i++ {
		j := rand.Intn(i + 1)
		Exchange(arr, i, j)
	}
}

// Zips two arrays together
// a := int[]{1, 2, 3}
// b := int[]{4, 5, 6}
// c, _ := Zip(&a, &b)
// c == int[]{1, 4, 2, 5, 3, 6}
func Zip[T any](arr1 *[]T, arr2 *[]T) ([]T, error) {
	if len(*arr1) != len(*arr2) {
		return nil, errors.New("arrays have different lengths")
	}
	result := make([]T, len(*arr1)+len(*arr2))
	for i := 0; i < len(*arr1)*2; i += 2 {
		result[i] = (*arr1)[i/2]
		result[i+1] = (*arr2)[i/2]
	}
	return result, nil
}

func Reverse[T any](arr *[]T) {
	for i := 0; i < len(*arr)/2; i++ {
		Exchange(arr, i, len(*arr)-i-1)
	}
}

func GetReversed[T any](arr *[]T) []T {
	acc := make([]T, 0, len(*arr))
	for i := len(*arr) - 1; i > 0; i-- {
		acc = append(acc, (*arr)[i])
	}
	return acc
}
