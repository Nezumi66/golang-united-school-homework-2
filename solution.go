package square

import (
	"math"
)

type intCustomType int

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	switch sidesNum {
	case 0:
		return math.Pi * math.Pow(sideLen, 2)
	case 3:
		return (math.Sqrt(3) * math.Pow(sideLen, 2)) / 4
	case 4:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
