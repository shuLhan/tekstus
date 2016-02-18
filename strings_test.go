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

var dStringsFindLongest = []struct {
	words []string
	exp   string
}{
	{
		[]string{"a", "bb", "ccc", "d", "eee"},
		"ccc",
	}, {
		[]string{"a", "bb", "ccc", "dddd", "eee"},
		"dddd",
	},
}

func TestStringsFindLongest(t *testing.T) {
	for _, td := range dStringsFindLongest {
		got, _ := tekstus.StringsFindLongest(td.words)

		assert(t, td.exp, got, true)
	}
}
