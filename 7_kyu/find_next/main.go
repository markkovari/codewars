package kata

import "math"

func FindNextSquare(sq int64) int64 {
	asFloat := float64(sq)
	floorSq := int64(math.Sqrt(asFloat))

	if floorSq*floorSq != sq {
		return -1
	}
	return (int64(floorSq) + 1) * int64((floorSq)+1)
}
