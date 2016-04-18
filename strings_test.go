// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"github.com/shuLhan/tekstus"
	"testing"
)

func TestStringsIsEqual(t *testing.T) {
	lss := tekstus.ListStrings{
		{"c", "b"},
		{"c", "a"},
	}
	check := []bool{true, false}

	for x, ss := range lss {
		b := tssExp[0][1].IsEqual(ss)

		if b != check[x] {
			t.Fatal(tssExp[0][1], " == ", ss, "? ", b)
		}
	}
}

func TestStringsIsContain(t *testing.T) {
	ss := tekstus.Strings{
		"a", "b", "c", "d",
	}
	ss2 := []string{
		"a", "b", "c", "d",
	}

	// Testing true positive
	got := tekstus.StringsIsContain(ss, "a")

	assert(t, true, got, true)

	// Testing true negative
	got = tekstus.StringsIsContain(ss, "e")

	assert(t, false, got, true)

	// Testing true positive
	got = tekstus.StringsIsContain(ss2, "a")

	assert(t, true, got, true)

	// Testing true negative
	got = tekstus.StringsIsContain(ss2, "e")

	assert(t, false, got, true)
}

func TestPartitioning(t *testing.T) {
	lss := tekstus.ListStrings{
		{"a", "b", "c"},
	}
	tssExp := []tekstus.TableStrings{{
		{{"a", "b", "c"}},
	}, {
		{{"a", "b"}, {"c"}},
		{{"b"}, {"a", "c"}},
		{{"a"}, {"b", "c"}},
	}, {
		{{"a"}, {"b"}, {"c"}},
	},
	}
	split := []int{1, 2, 3}

	for _, ss := range lss {
		for k := range split {
			setstr := ss.Partitioning(split[k])

			b := tssExp[k].IsEqual(setstr)
			if !b {
				t.Fatal("Expecting ", tssExp[k], " == ",
					setstr, "? ", b)
			}
		}
	}
}

func TestPartitioning2(t *testing.T) {
	lss := tekstus.ListStrings{
		{"a", "b", "c"},
		{"a", "b", "c", "d"},
		{"a", "b", "c", "d", "e"},
		{"a", "b", "c", "d", "e", "f"},
	}

	for _, ss := range lss {
		for i := 1; i <= len(ss); i++ {
			_ = ss.Partitioning(i)
		}
	}
}

func TestStringsSortByIndex(t *testing.T) {
	dat := []string{"Z", "X", "C", "V", "B", "N", "M"}
	exp := []string{"B", "C", "M", "N", "V", "X", "Z"}
	ids := []int{4, 2, 6, 5, 3, 1, 0}

	tekstus.StringsSortByIndex(&dat, ids)

	assert(t, exp, dat, true)
}
