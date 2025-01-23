package utils

import "math"

func CalculateColorDistance(r1, g1, b1, r2, g2, b2 int) float64 {
	return math.Pow(float64(r2-r1), 2) +
		math.Pow(float64(g2-g1), 2) +
		math.Pow(float64(b2-b1), 2)

}
