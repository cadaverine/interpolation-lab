package main

import "errors"

type Table struct {
	Xs []float64
	Ys []float64
}

func CreateTable(xArr []float64, yArr []float64) *Table {
	return &Table{
		Xs: xArr,
		Ys: yArr,
	}
}

type Config Table

func CreateConfig(xArr []float64, reducer func(float64) float64) *Config {
	yArr := make([]float64, len(xArr))

	for i, x := range xArr {
		yArr[i] = reducer(x)
	}

	return (*Config)(CreateTable(xArr, yArr))
}

func getNeighbors(values []float64, currentValue float64, pointsNum int) ([]float64, error) {
	valuesNum := len(values)

	if pointsNum > valuesNum {
		return nil, errors.New("length error")
	}

	if pointsNum == valuesNum {
		clone := make([]float64, len(values))
		copy(clone, values)

		return clone, nil
	}

	rightOffset := pointsNum / 2
	leftOffset := rightOffset + pointsNum%2

	var from, to int

	for i, value := range values {
		if value >= currentValue {
			if i < leftOffset {
				from = 0
				to = pointsNum - 1
			} else if valuesNum-1-i < rightOffset {
				to = valuesNum - 1
				from = to - (pointsNum - 1)
			} else {
				from = i - leftOffset
				to = i + rightOffset
			}

			break
		}
	}

	result := make([]float64, pointsNum)
	copy(result, values[from:to])

	return result, nil
}
