package main

import (
	"fmt"

	"github.com/sleepynut/GO-generic/sumutils"
)

func main() {
	floats := map[string]float64{
		"first":  40.2,
		"second": 29.8,
	}

	ints := map[string]int64{
		"first":  1,
		"second": 2,
	}

	fmt.Println("Sum floats: ", sumutils.SumMaps(floats))
	fmt.Println("Sum ints: ", sumutils.SumMaps(ints))
}
