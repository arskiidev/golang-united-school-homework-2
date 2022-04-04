package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type SideCount int

const (
	SidesTriangle SideCount = 3
	SidesSquare   SideCount = 4
	SidesCircle   SideCount = 0
)

func CalcSquare(sideLen float64, sidesNum SideCount) float64 {
	var result float64
	switch sidesNum {
	case 0:
		result = math.Pi * math.Pow(sideLen, 2)
	case 3:
		result = math.Sqrt(3) * math.Pow(sideLen, 2) / 4
	case 4:
		result = math.Pow(sideLen, 2)
	default:
		result = 0
	}
	return result
}
