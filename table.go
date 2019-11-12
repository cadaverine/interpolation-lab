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

func CreateTableFunction(xArr []float64, reducer func(float64) float64) *Table {
	yArr := make([]float64, len(xArr))

	for i, x := range xArr {
		yArr[i] = reducer(x)
	}

	return CreateTable(xArr, yArr)
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

func getNeighbors(values []float64, currentValue float64, pointsNum int) (int, int, error) {
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
