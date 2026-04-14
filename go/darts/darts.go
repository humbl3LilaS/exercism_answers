package darts

import (
	"math"
)

func Score(x, y float64) int {
	d := math.Sqrt((x*x) + (y*y))
	if d <= 1 {
		return 10
	}
	if d <= 5 {
		return 5
	}
	if d <= 10 {
		return 1
	}
	return 0
}
