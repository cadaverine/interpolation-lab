package main

import (
	"fmt"
	"math"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var reducers map[string]func(float64) float64 = map[string]func(float64) float64{
	"y = x^2":         func(x float64) float64 { return x * x },
	"y = x^3":         func(x float64) float64 { return x * x * x },
	"y = x^4":         func(x float64) float64 { return x * x * x * x },
	"y = cos(x)":      func(x float64) float64 { return math.Cos(x * math.Pi / 180) },
	"y = sqrt(x)":     func(x float64) float64 { return math.Sqrt(x) },
	"y = x * sqrt(x)": func(x float64) float64 { return x * math.Sqrt(x) },
	"y = x^2 - 1":     func(x float64) float64 { return x*x - 1 },
}

func main() {
	fmt.Print("Newton polynomial interpolation.\n\n")

	x0, h, x, n, N, reducerKey := handleInput(reducers)

	reducer := reducers[reducerKey]

	xArray := getRange(x0, h, n)
	yArray := getYArray(xArray, reducer)

	interpolated, _ := interpolate(xArray, yArray, N, x)
	root, _ := getRoots(xArray, yArray, N)

	printTable(xArray, yArray)

	fmt.Println("Function result: ", reducer(x))
	fmt.Println("Interpolation result: ", interpolated)
	fmt.Println("Function root: ", root)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Newton polynomial interpolation " + "(" + reducerKey + ")"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"Real function", getFunctionPlotPoints(xArray, yArray, reducer),
		"Interpolation", getInterpolatedPlotPoints(xArray, yArray, N),
	)

	if err != nil {
		panic(err)
	}

	if err := p.Save(5*vg.Inch, 10*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func handleInput(map[string]func(float64) float64) (x0, h, x float64, n, N int, reducerKey string) {
	fmt.Println("--------------------------")
	fmt.Print("SET TABLE FUNCTION PARAMS\n\n")

	fmt.Println("Choose the available function:")

	keys := make([]string, 0, len(reducers))
	for key := range reducers {
		keys = append(keys, key)
	}

	for i, key := range keys {
		fmt.Println(strconv.Itoa(i)+":", key)
	}

	var index int
	fmt.Scanln(&index)
	reducerKey = keys[index]

	fmt.Print("Enter x0: ")
	fmt.Scanln(&x0)

	fmt.Print("Enter h:  ")
	fmt.Scanln(&h)

	fmt.Print("Enter n:  ")
	fmt.Scanln(&n)

	fmt.Println("\n--------------------------")
	fmt.Print("SET INTERPOLATION PARAMS\n\n")

	fmt.Print("Enter N:  ")
	fmt.Scanln(&N)

	fmt.Print("Enter x:  ")
	fmt.Scanln(&x)

	return
}

func printTable(xArray, yArray []float64) {
	fmt.Printf("\n___________________\n")
	fmt.Printf("    X    |    Y    \n")
	for i := range xArray {
		fmt.Printf("%7.3f  | %7.3f\n", xArray[i], yArray[i])
	}
	fmt.Printf("-------------------\n\n")
}

func getDividedDifferences(xArray, yArray []float64) []float64 {
	differences := make([]float64, len(yArray))
	copy(differences, yArray)

	for i := 1; i < len(differences); i++ {
		previous := differences[i-1]
		for j := i; j < len(differences); j++ {
			temp := differences[j]
			differences[j] = (differences[j] - previous) / (xArray[j] - xArray[j-i])
			previous = temp
		}
	}

	return differences
}

func interpolate(xArray, yArray []float64, N int, x float64) (float64, error) {
	from, to, err := getNeighborsIndexes(xArray, x, N+1)

	if err != nil {
		return 0, err
	}

	xs := xArray[from : to+1]
	ys := yArray[from : to+1]

	diffs := getDividedDifferences(xs, ys)
	y := diffs[0]

	for i := 1; i < N+1; i++ {
		step := diffs[i]

		for j := 0; j < i; j++ {
			step *= x - xs[j]
		}

		y += step
	}

	return y, nil
}

func getRoots(xArray, yArray []float64, N int) (float64, error) {
	return interpolate(yArray, xArray, N, 0)
}

func getPoints(xArray, yArray []float64) plotter.XYs {
	points := make(plotter.XYs, len(xArray))

	for i, x := range xArray {
		points[i].X = x
		points[i].Y = yArray[i]
	}

	return points
}

func getFunctionPlotPoints(xArray, yArray []float64, reducer func(float64) float64) plotter.XYs {
	points := make(plotter.XYs, (len(xArray)-1)*10)

	for i, x := range xArray {
		if i != len(xArray)-1 {
			temp := xArray[i]
			step := (xArray[i+1] - xArray[i]) / 10
			for j := 0; j < 10; j++ {
				index := j + i*10
				xValue := temp + step*float64(j)
				points[index].X = xValue
				points[index].Y = reducer(xValue)
			}
		} else {
			points[10*i-1].X = x
			points[10*i-1].Y = reducer(x)
		}
	}

	return points
}

func getInterpolatedPlotPoints(xArray, yArray []float64, N int) plotter.XYs {
	points := make(plotter.XYs, (len(xArray)-1)*10)

	for i, x := range xArray {
		if i != len(xArray)-1 {
			temp := xArray[i]
			step := (xArray[i+1] - xArray[i]) / 10
			for j := 0; j < 10; j++ {
				index := j + i*10
				xValue := temp + step*float64(j)
				interpolated, _ := interpolate(xArray, yArray, N, xValue)

				points[index].X = xValue
				points[index].Y = interpolated
			}
		} else {
			interpolated, _ := interpolate(xArray, yArray, N, x)
			points[10*i-1].X = x
			points[10*i-1].Y = interpolated
		}
	}

	return points
}
