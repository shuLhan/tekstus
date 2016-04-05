// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"github.com/shuLhan/tekstus"
	"testing"
)

func TestFloat64Sum(t *testing.T) {
	in := []float64{1, 1.1, 1.2, 1.3, 1.4, -6}
	exp := 0.0
	got := tekstus.Float64Sum(in)

	assert(t, exp, got, true)
}

func TestFloat64Counts(t *testing.T) {
	data := []float64{1, 1, 2, 2, 3, 1, 2}
	classes := []float64{1, 2, 3}
	exp := []int{3, 3, 1}

	got := tekstus.Float64Counts(&data, &classes)

	assert(t, exp, got, true)
}
