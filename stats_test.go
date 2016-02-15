// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	_ "fmt"
	"github.com/shuLhan/tekstus"
	"testing"
)

func TestStringCountBy(t *testing.T) {
	data := []string{"A", "B", "A", "C"}
	class := []string{"A", "B"}
	exp := []int{2, 1}

	got := tekstus.StringCountBy(data, class)

	assert(t, exp, got, true)
}

var dataCharSequence = []struct {
	line     string
	nospace  bool
	expv     []rune
	expc     []int
	expvtest bool
	expctest bool
}{
	{
		"// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.",
		true, []rune{'/', 'l'}, []int{2, 2}, true, true,
	}, {
		"Use of this source code is governed by a BSD-style",
		true, nil, nil, true, true,
	}, {
		"aaa abcdee ffgf",
		true, []rune{'a', 'e', 'f'}, []int{3, 2, 2}, true, true,
	}, {
		" |  image name          = {{legend|#0080FF|Areas affected by flooding}}{{legend|#002255|Death(s) affected by flooding}}{{legend|#C83737|Areas affected by flooding and strong winds}}{{legend|#550000|Death(s) affected by flooding and strong winds}}",
		true,
		[]rune{'{', '0', 'F', 'f', 'o', '}', '{', '0', '2', '5', 'f', 'o', '}', '{', 'f', 'o', '}', '{', '5', '0', 'f', 'o', '}'},
		[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 4, 2, 2, 2},
		true, true,
	},
}

func doCountCharSequence(t *testing.T, line string, nospace bool, expv []rune,
	expc []int, expvtest, expctest bool,
) {
	gotv, gotc := tekstus.CountCharSequence(line, nospace)

	assert(t, expv, gotv, expvtest)
	assert(t, expc, gotc, expctest)
}

func TestCountCharSequence(t *testing.T) {
	for _, td := range dataCharSequence {
		doCountCharSequence(t, td.line, td.nospace, td.expv, td.expc,
			td.expvtest, td.expctest)
	}
}

var dataMaxCharSequence = []struct {
	char     rune
	count    int
	expvtest bool
	expctest bool
}{
	{'/', 2, true, true},
	{0, 0, true, true},
	{'a', 3, true, true},
	{'0', 4, true, true},
}

func doGetMaxCharSequence(t *testing.T, line string, char rune, count int,
	expvtest, expctest bool) {
	gotv, gotc := tekstus.GetMaxCharSequence(line, true)

	assert(t, char, gotv, expvtest)
	assert(t, count, gotc, expctest)
}

func TestGetMaxCharSequence(t *testing.T) {
	for x, td := range dataMaxCharSequence {
		doGetMaxCharSequence(t, dataCharSequence[x].line, td.char,
			td.count, td.expvtest, td.expctest)
	}
}

var dataUpperLowerTest = []struct {
	line  string
	expup int
	explo int
}{
	{"// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.", 4, 44},
}

func doCountUpperLowerChar(t *testing.T, line string, expup, explo int) {
	gotup, gotlo := tekstus.CountUpperLowerChar(line)

	assert(t, expup, gotup, true)
	assert(t, explo, gotlo, true)
}

func TestCountUpperLowerChar(t *testing.T) {
	for _, td := range dataUpperLowerTest {
		doCountUpperLowerChar(t, td.line, td.expup, td.explo)
	}
}

var dataCountDigit = []struct {
	line string
	exp  int
}{
	{"// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.", 4},
}

func doCountDigit(t *testing.T, line string, exp int) {
	got := tekstus.CountDigit(line)

	assert(t, exp, got, true)
}

func TestCountDigit(t *testing.T) {
	for _, td := range dataCountDigit {
		doCountDigit(t, td.line, td.exp)
	}
}

func TestCountAlnumChar(t *testing.T) {
	line := "// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved."
	expnon := 16

	n, l := tekstus.CountAlnumChar(line)

	assert(t, expnon, l-n, true)
}

func TestCountUniqChar(t *testing.T) {
	line := "// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved."
	exp := 34

	n, _ := tekstus.CountUniqChar(line)

	assert(t, exp, n, true)
}
