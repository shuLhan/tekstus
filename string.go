// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

import (
	"strings"
	"unicode"
)

/*
StringCountTokens given a text, count how many tokens inside of it and return
sum of all.
*/
func StringCountTokens(text string, tokens []string) (cnt int) {
	for _, v := range tokens {
		cnt += strings.Count(text, v)
	}
	return
}

/*
StringTrimNonAlnum remove non alpha-numeric character at the beginning and end
for `text`.
*/
func StringTrimNonAlnum(text string) string {
	r := []rune(text)
	rlen := len(r)
	start := 0

	for ; start < rlen; start++ {
		if unicode.IsLetter(r[start]) || unicode.IsDigit(r[start]) {
			break
		}
	}

	if start >= rlen {
		return ""
	}

	r = r[start:]
	rlen = len(r)
	end := rlen - 1
	for ; end >= 0; end-- {
		if unicode.IsLetter(r[end]) || unicode.IsDigit(r[end]) {
			break
		}
	}

	if end < 0 {
		return ""
	}

	r = r[:end+1]

	return string(r)
}

/*
StringUniq remove duplicate word from `words`. If sensitive is true then
compare the string with case sensitive.

Return the list of unique words.
*/
func StringUniq(words []string, sensitive bool) (uniques []string) {
	// Remove duplicate values.
	wordslen := len(words)
	xcmp := ""
	ycmp := ""

	for x, v := range words {
		if v == "" {
			continue
		}

		if sensitive {
			xcmp = v
		} else {
			xcmp = strings.ToLower(v)
		}

		for y := x + 1; y < wordslen; y++ {
			if len(words[y]) == 0 {
				continue
			}

			if sensitive {
				ycmp = words[y]
			} else {
				ycmp = strings.ToLower(words[y])
			}

			if xcmp == ycmp {
				words[y] = ""
			}
		}

		uniques = append(uniques, v)
	}
	return
}

/*
StringSplitWords given a text, return all words in text.

Definition of word are,
- any sequence of characters that is equal or greater than one that is
separated by space.
- does not start with number
- does not end with number

If cleanit is true remove any non-alphanumeric in the start and the end of
each words.

If uniq is true remove duplicate words.
*/
func StringSplitWords(text string, cleanit bool, uniq bool) (words []string) {
	words = strings.Fields(text)

	if !cleanit {
		return
	}

	// Clean the fields, remove non-alphanumeric character from start and
	// end.
	for x, word := range words {
		words[x] = StringTrimNonAlnum(word)
	}

	if !uniq {
		return
	}

	// Remove duplicate values.
	return StringUniq(words, false)
}
