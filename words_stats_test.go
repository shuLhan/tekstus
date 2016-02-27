// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"github.com/shuLhan/tekstus"
	"testing"
)

func TestWordsCountsOf(t *testing.T) {
	data := []string{"A", "B", "A", "C"}
	class := []string{"A", "B"}
	exp := []int{2, 1}

	got := tekstus.WordsCountTokens(data, class, false)

	assert(t, exp, got, true)
}

func TestWordsFrequenciesOf(t *testing.T) {
	words := []string{"a", "b", "a", "b", "a", "c"}
	tokens := []string{"a", "b"}
	wordslen := float64(len(words))
	exp := (3.0 / wordslen) + (2.0 / wordslen)

	got := tekstus.WordsFrequenciesOf(words, tokens, false)

	assert(t, exp, got, true)
}

func TestWordsCountMissRate(t *testing.T) {
	data := []string{"A", "B", "C", "D"}
	test := []string{"A", "B", "C", "D"}
	repl := []string{"B", "C", "D", "E"}
	exps := []float64{0.25, 0.5, 0.75, 1.0}

	got, _, _ := tekstus.WordsCountMissRate(data, test)
	assert(t, 0.0, got, true)

	for x, exp := range exps {
		test[x] = repl[x]
		got, _, _ = tekstus.WordsCountMissRate(data, test)
		assert(t, exp, got, true)
	}
}
