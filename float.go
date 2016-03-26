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
