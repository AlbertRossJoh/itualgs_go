package sorting

import "golang.org/x/exp/constraints"

func InsertionSort[T constraints.Ordered](arr *[]T) {
	n := len(*arr)

	exchanges := 0
	for i := n - 1; i > 0; i-- {
		if (*arr)[i] < (*arr)[i-1] {
			exchange(arr, i, i-1)
			exchanges++
		}
	}
	if exchanges == 0 {
		return
	}

	for i := 2; i < n; i++ {
		v := (*arr)[i]
		j := i
		for v < (*arr)[j-1] {
			(*arr)[j] = (*arr)[j-1]
			j--
		}
		(*arr)[j] = v
	}
}

func InsertionRangeSort[T constraints.Ordered](arr *[]T, start int, end int) {
	exchanges := 0
	for i := end - 1; i > start; i-- {
		if (*arr)[i] < (*arr)[i-1] {
			exchange(arr, i, i-1)
			exchanges++
		}
	}
	if exchanges == 0 {
		return
	}

	for i := start + 2; i < end; i++ {
		v := (*arr)[i]
		j := i
		for v < (*arr)[j-1] {
			(*arr)[j] = (*arr)[j-1]
			j--
		}
		(*arr)[j] = v
	}
}

func exchange[T any](arr *[]T, i int, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}
