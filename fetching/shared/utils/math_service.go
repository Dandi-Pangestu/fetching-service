package utils

func ConvertIDRTOUSD(idr float64, convertion float64) float64 {
	return idr / convertion
}

func CalcMedian(n int) float64 {
	if n%2 == 0 {
		return ((float64(n) / 2) + (float64(n)/2 + 1)) / 2
	} else {
		return (float64(n) + 1) / 2
	}
}
