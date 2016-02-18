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

var dataStringCountTokens = []struct {
	line   string
	tokens []string
	exp    int
}{
	{
		"// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.",
		[]string{"//"},
		1,
	}, {
		"The [[United States]] has regularly voted alone and against international consensus, using its [[United Nations Security Council veto power|veto power]] to block the adoption of proposed UN Security Council resolutions supporting the [[PLO]] and calling for a two-state solution to the [[Israeli-Palestinian conflict]].<ref>[http://books.google.ca/books?id=CHL5SwGvobQC&pg=PA168&dq=US+veto+Israel+regularly#v=onepage&q=US%20veto%20Israel%20regularly&f=false Pirates and emperors, old and new: international terrorism in the real world], [[Noam Chomsky]], p. 168.</ref><ref>The US has also used its veto to block resolutions that are critical of Israel.[http://books.google.ca/books?id=yzmpDAz7ZAwC&pg=PT251&dq=US+veto+Israel+regularly&lr=#v=onepage&q=US%20veto%20Israel%20regularly&f=false Uneasy neighbors], David T. Jones and David Kilgour, p. 235.</ref> The United States responded to the frequent criticism from UN organs by adopting the [[Negroponte doctrine]].",
		[]string{"[[", "]]", "<ref", "/ref>", "[http:"},
		18,
	},
}

func TestStringCountTokens(t *testing.T) {
	for _, td := range dataStringCountTokens {
		got := tekstus.StringCountTokens(td.line, td.tokens)

		assert(t, td.exp, got, true)
	}
}
