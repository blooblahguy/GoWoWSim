package core

import (
	"math"
	"math/rand"
)

func rage_factor(level int) float64 {
	rage_fac_1 := float64(0.0091107836)
	rage_fac_2 := float64(3.225598133)
	rage_fac_3 := float64(4.2652911)
	return rage_fac_1*float64(powint(level, 2)) + rage_fac_2*float64(level) + rage_fac_3
}

func powint(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func randint(min int, max int) int {
	// return int(min + rand.intn(max-min))
	// rand.Seed(time.Now().UnixNano())
	return (rand.Intn(max-min+1) + min)
}

func minint(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxint(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
