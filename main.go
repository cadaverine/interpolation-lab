package main

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	fmt.Println("Newton polynomial interpolation (y = x^2).")

	// var x0, h, n string

	// fmt.Print("Enter x0: ")
	// fmt.Scanln(&x0)

	// fmt.Print("Enter h:  ")
	// fmt.Scanln(&h)

	// fmt.Print("Enter n:  ")
	// fmt.Scanln(&n)

	xArray := getRange(2.3, 0.6, 10)
	yArray := getYArray(xArray, func(x float64) float64 { return x * x })

	fmt.Println("xArray: ", xArray)
	fmt.Println("yArray: ", yArray)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Newton polynomial interpolation (y = x^2)"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p, "Base", getPoints(xArray, yArray))

	if err != nil {
		panic(err)
	}

	if err := p.Save(5*vg.Inch, 10*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func getPoints(xArray, yArray []float64) plotter.XYs {
	points := make(plotter.XYs, len(xArray))

	for i, x := range xArray {
		points[i].X = x
		points[i].Y = yArray[i]
	}

	return points
}
