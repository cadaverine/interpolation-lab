package main

import (
	"errors"
	"math"
)

func getRange(from, step float64, pointsNum int) []float64 {
	rangeArr := make([]float64, pointsNum)

	rangeArr[0] = from
	for i := 1; i < pointsNum; i++ {
		rangeArr[i] = rangeArr[i-1] + step
	}

	return rangeArr
}

func getYArray(xArray []float64, reducer func(float64) float64) []float64 {
	yArray := make([]float64, len(xArray))

	for i, x := range xArray {
		yArray[i] = reducer(x)
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

func findIndex(values []float64, currentValue float64) (int, error) {
	var err error
	var result int

	switch true {
	case len(values) == 0:
		err = errors.New("Error: values slice is empty")

	case currentValue > values[len(values)-1]:
		result = len(values) - 1

	default:
		for i, value := range values {
			if value >= currentValue {
				result = i
				break
			}
		}
	}

	return result, err
}

func getNeighborsIndexes(values []float64, currentValue float64, pointsNum int) (int, int, error) {
	var err error
	var from, to int

	valuesNum := len(values)

	switch true {
	case pointsNum < 2:
		err = errors.New("Error: points num must be more than 1")

	case pointsNum > valuesNum:
		err = errors.New("Error: values num less then points num")

	case pointsNum == valuesNum:
		to = len(values) - 1

	default:
		rightOffset := (pointsNum - 1) / 2
		leftOffset := rightOffset + (pointsNum-1)%2
		index, err := findIndex(values, currentValue)

		if err == nil {
			if index < leftOffset {
				to = pointsNum - 1
			} else if valuesNum-1 < index+rightOffset {
				from = valuesNum - 1 - pointsNum
				to = valuesNum - 1
			} else {
				from = index - leftOffset
				to = index + rightOffset
			}
		}
	}

	return from, to, err
}
