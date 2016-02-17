// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"github.com/shuLhan/tekstus"
	"testing"
)

func TestIntFindMax(t *testing.T) {
	in1 := []int{}
	in2 := []int{1, 2, 3, 4, 5}

	maxv, maxid := tekstus.IntFindMax(in1)

	assert(t, -1, maxid, true)

	maxv, maxid = tekstus.IntFindMax(in2)

	assert(t, 5, maxv, true)
	assert(t, 4, maxid, true)
}

func TestIntFindMin(t *testing.T) {
	in1 := []int{}
	in2 := []int{1, 2, 3, 4, 5}

	minv, minid := tekstus.IntFindMin(in1)

	assert(t, -1, minid, true)

	minv, minid = tekstus.IntFindMin(in2)

	assert(t, 1, minv, true)
	assert(t, 0, minid, true)
}

func TestIntSum(t *testing.T) {
	data := []int{0, 1, 2, 3, -3, -2, -1, 0}
	exp := 0

	got := tekstus.IntSum(data)

	assert(t, exp, got, true)
}
