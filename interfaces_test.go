// Copyright 2016-2018 Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"github.com/shuLhan/tekstus"
	"testing"
)

func TestInterfacesToStrings(t *testing.T) {
	is := make([]interface{}, 0)
	i64 := []int64{0, 1, 2, 3}
	exp := []string{"0", "1", "2", "3"}

	for _, v := range i64 {
		is = append(is, v)
	}

	got := tekstus.InterfacesToStrings(is)

	assert(t, exp, got, true)
}

func TestInterfacesToStrings2(t *testing.T) {
	is := make([]interface{}, 0)
	f64 := []float64{0, 1.1, 2.2, 3.3}
	exp := []string{"0", "1.1", "2.2", "3.3"}

	for _, v := range f64 {
		is = append(is, v)
	}

	got := tekstus.InterfacesToStrings(is)

	assert(t, exp, got, true)
}
