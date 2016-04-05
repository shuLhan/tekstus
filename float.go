// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

//
// Float64Sum return sum of slice of float64.
//
func Float64Sum(data []float64) (sum float64) {
	for _, v := range data {
		sum += v
	}
	return
}

//
// Float64Count will count number of class in data.
//
func Float64Count(data *[]float64, class float64) (count int) {
	if len(*data) <= 0 {
		return
	}

	for _, v := range *data {
		if v == class {
			count++
		}
	}
	return
}

//
// Float64Counts will count class in data and return each of the counter.
//
// For example, if data is "[1,1,2]" and class is "[1,2]", this function will
// return "[2,1]".
//
//	idx class  count
//	0 : 1   -> 2
//	1 : 2   -> 1
//
//
func Float64Counts(data *[]float64, classes *[]float64) (counts []int) {
	if len(*classes) <= 0 {
		return
	}

	counts = make([]int, len(*classes))

	for x, c := range *classes {
		counts[x] = Float64Count(data, c)
	}
	return
}
