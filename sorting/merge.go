package sorting

import (
	"sync"

	"golang.org/x/exp/constraints"
)

var max = 1 << 11

func MergeSort[T constraints.Ordered](arr *[]T) {
	var aux []T
	for _, elm := range *arr {
		aux = append(aux, elm)
	}
	sort_parallel(&aux, arr, 0, len(*arr)-1)
}

func merge[T constraints.Ordered](src *[]T, dst *[]T, lo int, mid int, hi int) {
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			(*dst)[k] = (*src)[j]
			j++
		} else if j > hi {
			(*dst)[k] = (*src)[i]
			i++
		} else if (*src)[i] > (*src)[j] {
			(*dst)[k] = (*src)[j]
			j++
		} else {
			(*dst)[k] = (*src)[i]
			i++
		}
	}
}

func sort[T constraints.Ordered](src *[]T, dst *[]T, lo int, hi int) {
	if hi <= lo+7 {
		insertionSort(dst, lo, hi)
		return
	}

	mid := lo + (hi-lo)/2
	sort(dst, src, lo, mid)
	sort(dst, src, mid+1, hi)

	if !((*src)[mid+1] < (*src)[mid]) {
		for i := lo; i <= hi; i++ {
			(*dst)[i] = (*src)[i]
		}
		return
	}

	merge(src, dst, lo, mid, hi)
}

func sort_parallel[T constraints.Ordered](src *[]T, dst *[]T, lo int, hi int) {
	if len(*src) <= max {
		sort(src, dst, lo, hi)
		return
	}
	if hi <= lo+7 {
		insertionSort(dst, lo, hi)
		return
	}

	mid := lo + (hi-lo)/2

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		sort_parallel(dst, src, lo, mid)
	}()

	sort_parallel(dst, src, mid+1, hi)

	wg.Wait()
	if !((*src)[mid+1] < (*src)[mid]) {
		for i := lo; i <= hi; i++ {
			(*dst)[i] = (*src)[i]
		}
		return
	}

	merge(src, dst, lo, mid, hi)
}

func insertionSort[T constraints.Ordered](arr *[]T, lo int, hi int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && (*arr)[j] < (*arr)[j-1]; j-- {
			(*arr)[j], (*arr)[j-1] = (*arr)[j-1], (*arr)[j]
		}
	}
}

func copyRange[T constraints.Ordered](src *[]T, dst *[]T, lo int, hi int) {
	for i := lo; i <= hi; i++ {
		(*src)[i] = (*dst)[i]
	}
}
