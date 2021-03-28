package movingaverage_test

import (
	"fmt"
	"github.com/SimonWaldherr/golang-moving-average"
)

func ExampleMovingAverage() {
	values := movingaverage.New(4)

	values.Add(1)
	fmt.Printf("Values: %#v\nSum: %.0f, Arithmetic: %.1f, Median: %.1f, Geometric: %.2f\n\n",
		values.FilledValues(),
		values.Sum(),
		values.Arithmetic(),
		values.Median(),
		values.Geometric(),
	)

	values.Add(2)
	fmt.Printf("Values: %#v\nSum: %.0f, Arithmetic: %.1f, Median: %.1f, Geometric: %.2f\n\n",
		values.FilledValues(),
		values.Sum(),
		values.Arithmetic(),
		values.Median(),
		values.Geometric(),
	)

	values.Add(6)
	fmt.Printf("Values: %#v\nSum: %.0f, Arithmetic: %.1f, Median: %.1f, Geometric: %.2f\n\n",
		values.FilledValues(),
		values.Sum(),
		values.Arithmetic(),
		values.Median(),
		values.Geometric(),
	)

	values.Add(3)
	fmt.Printf("Values: %#v\nSum: %.0f, Arithmetic: %.1f, Median: %.1f, Geometric: %.2f\n\n",
		values.FilledValues(),
		values.Sum(),
		values.Arithmetic(),
		values.Median(),
		values.Geometric(),
	)

	values.Add(0)
	fmt.Printf("Values: %#v\nSum: %.0f, Arithmetic: %.1f, Median: %.1f, Geometric: %.2f\n\n",
		values.FilledValues(),
		values.Sum(),
		values.Arithmetic(),
		values.Median(),
		values.Geometric(),
	)

	// Output:
	// Values: []float64{1}
	// Sum: 1, Arithmetic: 1.0, Median: 1.0, Geometric: 1.00
	//
	// Values: []float64{1, 2}
	// Sum: 3, Arithmetic: 1.5, Median: 1.5, Geometric: 1.41
	//
	// Values: []float64{1, 2, 6}
	// Sum: 9, Arithmetic: 3.0, Median: 2.0, Geometric: 2.29
	//
	// Values: []float64{1, 2, 3, 6}
	// Sum: 12, Arithmetic: 3.0, Median: 2.5, Geometric: 2.45
	//
	// Values: []float64{0, 2, 3, 6}
	// Sum: 11, Arithmetic: 2.8, Median: 2.5, Geometric: 0.00
}
