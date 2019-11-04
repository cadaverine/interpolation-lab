package main

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
