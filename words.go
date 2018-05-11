// Copyright 2016-2018 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

import (
	"strings"
)

//
// WordsUniq remove duplicate word from `words`. If sensitive is true then
// compare the string with case sensitive.
//
// Return the list of unique words.
//
func WordsUniq(words []string, sensitive bool) (uniques []string) {
	// Remove duplicate values.
	wordslen := len(words)
	var xcmp, ycmp string

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

//
// WordsFindLongest find the longest word in words and return their value and
// index.
//
// If words is empty return nil string with negative (-1) index.
//
func WordsFindLongest(words []string) (slong string, idx int) {
	if len(words) <= 0 {
		return "", -1
	}

	slonglen := len(slong)

	for x, v := range words {
		vlen := len(v)
		if vlen > slonglen {
			slonglen = vlen
			slong = v
			idx = x
		}
	}
	return
}
