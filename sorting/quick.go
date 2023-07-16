package sorting

import (
	utils "github.com/AlbertRossJoh/itualgs_go/sharedfunctions"
	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Ordered](arr *[]T) {
	utils.Shuffle(arr)
	quickSort(arr, 0, len(*arr)-1)
}

func quickSort[T constraints.Ordered](arr *[]T, low, high int) {
	if high <= low {
		return
	}
	j := partition(arr, low, high)
	quickSort(arr, low, j-1)
	quickSort(arr, j+1, high)
}

func QuickSelect[T constraints.Ordered](arr *[]T, k int) T {
	utils.Shuffle(arr)
	lo, hi := 0, len(*arr)-1
	for lo < hi {
		i := partition(arr, lo, hi)
		if i > k {
			hi = i - 1
		} else if i < k {
			lo = i + 1
		} else {
			return (*arr)[i]
		}
	}
	return (*arr)[lo]
}

func partition[T constraints.Ordered](arr *[]T, low, high int) int {
	i, j, v := low, high+1, (*arr)[low]

	i++
	j--
	for {
		for (*arr)[i] < v {
			if i == high {
				break
			}
			i++
		}
		for (*arr)[j] > v {
			if j == low {
				break
			}
			j--
		}
		if i >= j {
			break
		}

		utils.Exchange(arr, i, j)
	}

	utils.Exchange(arr, low, j)

	return j
}
