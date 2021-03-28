package movingaverage

import (
	"errors"
	"math"
	"sort"
)

// @author Robin Verlangen
// Moving average implementation for Go

var errNoValues = errors.New("no values")

type MovingAverage struct {
	Window          int
	values          []float64
	valPos          int
	slotsFilled     bool
	ignoreNanValues bool
	ignoreInfValues bool
}

func (ma *MovingAverage) SetIgnoreInfValues(ignoreInfValues bool) {
	ma.ignoreInfValues = ignoreInfValues
}

func (ma *MovingAverage) SetIgnoreNanValues(ignoreNanValues bool) {
	ma.ignoreNanValues = ignoreNanValues
}

func (ma *MovingAverage) Arithmetic() float64 {
	var sum = float64(0)
	values := ma.FilledValues()
	if values == nil {
		return 0
	}
	n := len(values)
	for _, value := range values {
		sum += value
	}

	avg := sum / float64(n)
	return avg
}

func (ma *MovingAverage) Avg() float64 {
	return ma.Arithmetic()
}

func (ma *MovingAverage) Median() float64 {
	values := ma.FilledValues()
	c := len(values)

	sort.Float64s(values)
	if c%2 == 1 {
		return values[c/2]
	}
	return (values[c/2] + values[c/2-1]) / 2
}

func (ma *MovingAverage) Geometric() float64 {
	values := ma.FilledValues()
	c := len(values)

	var m float64 = 1
	for _, value := range values {
		m = m * value
	}
	return math.Pow(m, 1/float64(c))
}

func (ma *MovingAverage) FilledValues() []float64 {
	var c = ma.Window - 1

	// Are all slots filled? If not, ignore unused
	if !ma.slotsFilled {
		c = ma.valPos - 1
		if c < 0 {
			// Empty register
			return nil
		}
	}
	return ma.values[0 : c+1]
}

func (ma *MovingAverage) Add(values ...float64) {
	for _, val := range values {
		// ignore NaN?
		if ma.ignoreNanValues && math.IsNaN(val) {
			continue
		}

		// ignore Inf?
		if ma.ignoreInfValues && math.IsInf(val, 0) {
			continue
		}

		// Put into values array
		ma.values[ma.valPos] = val

		// Increment value position
		ma.valPos = (ma.valPos + 1) % ma.Window

		// Did we just go back to 0, effectively meaning we filled all registers?
		if !ma.slotsFilled && ma.valPos == 0 {
			ma.slotsFilled = true
		}
	}
}

func (ma *MovingAverage) SlotsFilled() bool {
	return ma.slotsFilled
}

func (ma *MovingAverage) Values() []float64 {
	return ma.FilledValues()
}

func (ma *MovingAverage) Count() int {
	return len(ma.Values())
}

func (ma *MovingAverage) Sum() float64 {
	var sum float64
	values := ma.FilledValues()
	if values == nil {
		return 0
	}

	for _, value := range values {
		sum += value
	}
	return sum
}

func (ma *MovingAverage) Max() (float64, error) {
	best := math.MaxFloat64 * -1
	values := ma.FilledValues()
	if values == nil {
		return 0, errNoValues
	}
	for _, value := range values {
		if value > best {
			best = value
		}
	}
	return best, nil
}

func (ma *MovingAverage) Min() (float64, error) {
	if !ma.slotsFilled && ma.valPos == 0 {
		return 0, errNoValues
	}
	best := math.MaxFloat64
	values := ma.FilledValues()
	for _, value := range values {
		if value < best {
			best = value
		}
	}
	return best, nil
}

func New(window int) *MovingAverage {
	return &MovingAverage{
		Window:      window,
		values:      make([]float64, window),
		valPos:      0,
		slotsFilled: false,
	}
}
