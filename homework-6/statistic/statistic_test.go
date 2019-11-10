package statistic

import "testing"

type testpair struct {
	values []float64
	result float64
}

var testsAverage = []testpair {
	{[]float64{1,2}, 1.5},
	{[]float64{1,1,1,1,1,1}, 1},
	{[]float64{-1,1}, 0},
}

var testsSum = []testpair {
	{[]float64{1,2}, 3},
	{[]float64{1,1,1,1,1,1}, 6},
	{[]float64{-1,1}, 0},
	{[]float64{0,-1}, -1},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range testsAverage {
		v := Average(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values, 
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func TestSumSet(t *testing.T) {
	for _, pair := range testsSum {
		v := Sum(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values, 
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
