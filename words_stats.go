// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

import (
	"strings"
)

/*
WordsCountToken will return number of token occurence in words.
*/
func WordsCountToken(words []string, token string, sensitive bool) (cnt int) {

	if !sensitive {
		token = strings.ToLower(token)
	}

	for _, v := range words {
		if !sensitive {
			v = strings.ToLower(v)
		}

		if v == token {
			cnt++
		}
	}
	return
}

/*
WordsCountTokens count number of occurrence of each `tokens` values in words.
Return number of each tokens based on their index.

For example, if words is "[A,A,B]" and tokens is "[A,B]", this function will
return "[2,1]".

	idx cls  count
	0 : A -> 2
	1 : B -> 1
*/
func WordsCountTokens(words []string, tokens []string, sensitive bool) (
	clsCnt []int,
) {
	tokenslen := len(tokens)
	if tokenslen <= 0 {
		return
	}

	clsCnt = make([]int, tokenslen)

	for k, v := range tokens {
		clsCnt[k] = WordsCountToken(words, v, sensitive)
	}

	return
}

/*
WordsFrequencyOf return frequency of token in words using

	count-of-token / total-words
*/
func WordsFrequencyOf(words []string, token string, sensitive bool) float64 {
	wordslen := float64(len(words))
	if wordslen <= 0 {
		return 0
	}

	cnt := WordsCountToken(words, token, sensitive)

	return float64(cnt) / wordslen
}

/*
WordsFrequenciesOf return total frequency of tokens in words.
*/
func WordsFrequenciesOf(words []string, tokens []string, sensitive bool) (
	sumfreq float64,
) {
	if len(words) <= 0 || len(tokens) <= 0 {
		return 0
	}

	for _, token := range tokens {
		sumfreq += WordsFrequencyOf(words, token, sensitive)
	}
	return
}

/*
WordsMaxCountOf return the string that has highest frequency.

Example, given input

	words:  [A A B A B C C]
	tokens: [A B]

it will return A as the majority tokens in words.
If tokens has equal frequency, then the first tokens in order will returned.
*/
func WordsMaxCountOf(words []string, tokens []string, sensitive bool) string {
	if len(words) <= 0 {
		return ""
	}

	tokensCount := WordsCountTokens(words, tokens, sensitive)
	_, maxIdx := IntFindMax(tokensCount)

	if maxIdx < 0 {
		return ""
	}

	return tokens[maxIdx]
}
