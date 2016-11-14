package numgo

import (
	"math"
)

func Deg2Rad(num float64) float64 {
	return num * math.Pi / 180
}

func Rad2Deg(num float64) float64 {
	return num * 180 / math.Pi
}
