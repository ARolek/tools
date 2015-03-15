package tools

import (
	"math"
)

//	rounds a float to the nearest integer
func RoundFloat(float float64) int64 {
	return int64(math.Floor(float + 0.5))
}
