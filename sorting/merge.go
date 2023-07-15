package sorting

import (
	"golang.org/x/exp/constraints"
)

func MergeSort[T constraints.Ordered](arr *[]T) {
	aux := *arr
	sort(arr, &aux, 0, len(*arr)-1)
}

func merge[T constraints.Ordered](src *[]T, dst *[]T, lo int, mid int, hi int) {
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			(*dst)[k] = (*src)[j]
			j++
		} else if j > hi {
			(*dst)[k] = (*src)[i]
			i++
		} else if (*src)[i] < (*src)[j] {
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
		InsertionRangeSort(src, lo, hi)
		return
	}

	mid := lo + (hi-lo)/2
	sort(src, dst, lo, mid)
	sort(src, dst, mid+1, hi)

	if (*src)[mid] <= (*src)[mid+1] {
		for i := lo; i <= hi; i++ {
			(*dst)[i] = (*src)[i]
		}
		return
	}

	merge(src, dst, lo, mid, hi)
}
