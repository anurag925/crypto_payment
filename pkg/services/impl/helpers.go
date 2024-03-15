package impl

import "math"

func convertTo2Decimal(amount float64) float64 {
	return math.Round(amount*100) / 100
}
