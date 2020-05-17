package utils

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestConvertIDRTOUSD(t *testing.T) {
	testCases := []struct {
		idr        float64
		convertion float64
		expected   float64
	}{
		{74000, 14500, 5.103448275862069},
		{58000, 14100.5, 4.113329314563313},
	}

	for _, tc := range testCases {
		result := ConvertIDRTOUSD(tc.idr, tc.convertion)
		assert.Equal(t, result, tc.expected)
	}
}

func TestCalcMedian(t *testing.T) {
	testCases := []struct {
		n        int
		expected float64
	}{
		{79, 40},
		{70, 35.5},
	}

	for _, tc := range testCases {
		result := CalcMedian(tc.n)
		assert.Equal(t, result, tc.expected)
	}
}
