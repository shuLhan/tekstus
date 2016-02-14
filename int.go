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
