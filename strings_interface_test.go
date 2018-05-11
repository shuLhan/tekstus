// Copyright 2016-2018 Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"github.com/shuLhan/tekstus"
	"testing"
)

func TestStringsToInt64(t *testing.T) {
	in := []string{"0", "1", "e", "3.3"}
	exp := []int64{0, 1, 0, 3}

	got := tekstus.StringsToInt64(in)

	assert(t, exp, got, true)
}

func TestStringsToFloat64(t *testing.T) {
	in := []string{"0", "1.1", "e", "3"}
	exp := []float64{0, 1.1, 0, 3}

	got := tekstus.StringsToFloat64(in)

	assert(t, exp, got, true)
}
