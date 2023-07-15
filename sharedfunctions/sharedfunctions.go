package sharedfunctions

import "math"

func IsClose(x float64, y float64) bool {
	return math.Abs(x-y) < 0.000001
}
