package utils

import (
	"math"
)

// Haversine formula to calculate the distance between two points on the earth
func CalculateDistance(lat1, lng1, lat2, lng2 float64) float64 {
	var R = 6371e3                // metres
	var φ1 = lat1 * math.Pi / 180 // φ, λ in radians
	var φ2 = lat2 * math.Pi / 180
	var Δφ = (lat2 - lat1) * math.Pi / 180
	var Δλ = (lng2 - lng1) * math.Pi / 180

	var a = math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	var distance = R * c // in metres
	return distance
}
