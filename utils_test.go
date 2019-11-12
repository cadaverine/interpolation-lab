package main

import (
	"testing"
)

type inputData struct {
	values       []float64
	currentValue float64
	pointsNum    int
}

type outputData struct {
	from int
	to   int
	err  bool
}

type TestCase struct {
	input  *inputData
	output *outputData
}

func generateTestCases() []*TestCase {
	testCases := []*TestCase{
		&TestCase{
			&inputData{[]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}, 1.5, 5},
			&outputData{0, 4, false},
		},
		&TestCase{
			&inputData{[]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}, 6.9, 2},
			&outputData{5, 6, false},
		},
		&TestCase{
			&inputData{[]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}, 4.1, 3},
			&outputData{3, 5, false},
		},
		&TestCase{
			&inputData{[]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}, 4.1, 4},
			&outputData{2, 5, false},
		},
		&TestCase{
			&inputData{[]float64{1.0, 2.0, 3.0}, 0.0, 2},
			&outputData{0, 1, false},
		},
		&TestCase{
			&inputData{[]float64{1.0, 2.0, 3.0}, 3.1, 1},
			&outputData{0, 0, true},
		},
		&TestCase{
			&inputData{[]float64{1.0, 2.0, 3.0}, 3.1, 4},
			&outputData{0, 0, true},
		},
	}

	return testCases
}

func Test_getNeighborsIndexes(t *testing.T) {
	testCases := generateTestCases()

	for _, c := range testCases {
		from, to, err := getNeighborsIndexes(c.input.values, c.input.currentValue, c.input.pointsNum)

		if from != c.output.from || to != c.output.to {
			t.Errorf("Error: results is not equal.\nResult from: %v, to: %v\nExpected from: %v, to: %v", from, to, c.output.from, c.output.to)
		}

		if err == nil && c.output.err {
			t.Errorf("Error: error is expected.")
		} else if err != nil && !c.output.err {
			t.Errorf("Error: error is not expected.")
		}
	}
}
