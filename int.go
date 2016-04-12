// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

/*
IntFindMax given slice of integer, return the maximum value in slice and index
of maximum value.
If data is empty, return -1 in value and index.

Example, given a slice of data: [0 1 2 3 4], it will return 4 as max and 4 as
index of maximum value.
*/
func IntFindMax(data []int) (max int, maxidx int) {
	l := len(data)
	if l <= 0 {
		return -1, -1
	}

	i := 0
	max = data[i]
	maxidx = i

	for i = 1; i < l; i++ {
		if data[i] > max {
			max = data[i]
			maxidx = i
		}
	}

	return
}

//
// Int64FindMax given slice of integer, return the maximum value in slice and
// index of maximum value.
// If data is empty, return -1 in value and index.
//
func Int64FindMax(data []int64) (max int64, maxidx int) {
	l := len(data)
	if l <= 0 {
		return -1, -1
	}

	i := 0
	max = data[i]
	maxidx = i

	for i = 1; i < l; i++ {
		if data[i] > max {
			max = data[i]
			maxidx = i
		}
	}

	return
}

/*
IntFindMin given slice of integer, return the minimum value in slice and index
of minimum value.
If data is empty, return -1 in value and index.

Example, given a slice of data: [0 1 2 3 4], it will return 0 as min and 0 as
minimum index.
*/
func IntFindMin(data []int) (min int, minidx int) {
	l := len(data)
	if l <= 0 {
		return -1, -1
	}

	i := 0
	min = data[i]
	minidx = i

	for i = 1; i < l; i++ {
		if data[i] < min {
			min = data[i]
			minidx = i
		}
	}

	return
}

/*
IntSum return sum of all value in slice.
*/
func IntSum(data []int) (sum int) {
	for _, v := range data {
		sum += v
	}
	return
}

//
// Int64Sum return sum of all integer value.
//
func Int64Sum(data []int64) (sum int64) {
	for _, v := range data {
		sum += v
	}
	return
}

//
// Int64MaxCountOf will count number of occurence of each element of classes in
// data and return the class with maximum count.
// For example, given a data [0, 1, 0, 1, 0] and classes [0, 1], the counter
// will count 0 as 3, 1 as 2; and return 0.
//
func Int64MaxCountOf(data, classes []int64) int64 {
	clsCount := make([]int, len(classes))

	for x, c := range classes {
		v := 0
		for _, d := range data {
			if c == d {
				v++
			}
		}
		clsCount[x] = v
	}

	_, midx := IntFindMax(clsCount)

	return classes[midx]
}
