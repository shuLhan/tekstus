// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"github.com/shuLhan/tekstus"
	"testing"
)

var dataTestRunesDiff = []struct {
	l   []rune
	r   []rune
	exp []rune
}{
	{
		[]rune{'a', 'b', 'a', 'b', 'c', 'd'},
		[]rune{'d', 'c', 'a', 'e'},
		[]rune{'b', 'e'},
	}, {
		[]rune{'a', 'b', 'a', 'b', 'c', 'd'},
		[]rune{'d', 'c', 'a', 'e'},
		[]rune{'b', 'e'},
	}, {
		[]rune{'a', 'b', 'a', 'b', 'c', 'd'},
		[]rune{'d', 'c', 'a', 'b', 'a', 'b', 'e'},
		[]rune{'e'},
	}, {
		[]rune{'d', 'c', 'a', 'b', 'a', 'b', 'e'},
		[]rune{'a', 'b', 'f', 'a', 'b', 'c', 'd'},
		[]rune{'e', 'f'},
	},
}

func doRunesDiff(t *testing.T, l, r, exp []rune) {
	got := tekstus.RunesDiff(l, r)

	assert(t, string(exp), string(got), true)
}

func TestRunesDiff(t *testing.T) {
	for _, td := range dataTestRunesDiff {
		doRunesDiff(t, td.l, td.r, td.exp)
	}
}

func TestRunesEncapsulateTrim(t *testing.T) {
	for _, td := range dataEncapsulateTrim {
		got, _ := tekstus.RunesEncapsulateTrim([]rune(td.text),
			[]rune(td.leftcap), []rune(td.rightcap))

		assert(t, string(td.exp), string(got), true)
	}
}
