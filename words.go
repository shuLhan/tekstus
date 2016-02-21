// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

import (
	"strings"
)

/*
WordsUniq remove duplicate word from `words`. If sensitive is true then
compare the string with case sensitive.

Return the list of unique words.
*/
func WordsUniq(words []string, sensitive bool) (uniques []string) {
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
