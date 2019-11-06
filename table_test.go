package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	values         []float64
	currentValue   float64
	pointsNum      int
	expectedResult []float64
	hasError       bool
}

func generateTestCases() []*TestCase {
	testCases := []*TestCase{
		&TestCase{[]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}, 1.5, 5, []float64{1.0, 2.0, 3.0, 4.0, 5.0}, false},
		&TestCase{[]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}, 6.9, 2, []float64{6.0, 7.0}, false},
	}

	return testCases
}

func Test_getNeighbors(t *testing.T) {
	testCases := generateTestCases()

	for _, c := range testCases {
		result, err := getNeighbors(c.values, c.currentValue, c.pointsNum)

		if !reflect.DeepEqual(result, c.expectedResult) {
			t.Errorf("Error: results is not equal. Result: %v, Expected: %v", result, c.expectedResult)
		}

		if err == nil && c.hasError {
			t.Errorf("Error: error is expected.")
		} else if err != nil && !c.hasError {
			t.Errorf("Error: error is not expected.")
		}
	}
}
