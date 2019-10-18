package main

import "math"

func getRange(from, step float64, pointsNum int) []float64 {
	rangeArr := make([]float64, pointsNum)

	rangeArr[0] = from
	for i := 1; i < pointsNum; i++ {
		rangeArr[i] = round(rangeArr[i-1]+step, 4)
	}

	return rangeArr
}

func getYArray(xArray []float64, reducer func(float64) float64) []float64 {
	yArray := make([]float64, len(xArray))

	for i, x := range xArray {
		yArray[i] = round(reducer(x), 4)
	}

	return yArray
}

func round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
