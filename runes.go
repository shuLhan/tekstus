// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

import (
	"unicode"
)

/*
RunesContain return true if character `c` is in slice of rune `s` and index of
character in `s`.
*/
func RunesContain(s []rune, c rune) (bool, int) {
	for x, v := range s {
		if v == c {
			return true, x
		}
	}
	return false, -1
}

/*
RunesDiff return the difference between slice l and r.
*/
func RunesDiff(l []rune, r []rune) (diff []rune) {
	found := false
	dupDiff := []rune{}

	// Find l not in r
	for _, v := range l {
		found, _ = RunesContain(r, v)
		if !found {
			dupDiff = append(dupDiff, v)
		}
	}

	// Find r not in diff
	for _, v := range r {
		found, _ = RunesContain(l, v)
		if !found {
			dupDiff = append(dupDiff, v)
		}
	}

	// Remove duplicate in dupDiff
	duplen := len(dupDiff)
	for x, v := range dupDiff {
		found = false
		for y := x + 1; y < duplen; y++ {
			if v == dupDiff[y] {
				found = true
				break
			}
		}
		if !found {
			diff = append(diff, v)
		}
	}

	return
}

/*
RunesFindToken will search token in text starting from index `startAt` and
return the matching index.
*/
func RunesFindToken(line, token []rune, startAt int) (at int) {
	y := 0
	tokenlen := len(token)
	linelen := len(line)

	at = -1
	for x := startAt; x < linelen; x++ {
		if line[x] == token[y] {
			if y == 0 {
				at = x
			}
			y++
			if y == tokenlen {
				// we found it!
				return
			}
		} else {
			if at != -1 {
				// reset back
				y = 0
				at = -1
			}
		}
	}
	// x run out before y
	if y < tokenlen {
		at = -1
	}
	return
}

/*
RunesFindSpaces in line, return -1 if not found.
*/
func RunesFindSpaces(line []rune, startAt int) (idx int) {
	lineLen := len(line)

	for idx = startAt; idx < lineLen; idx++ {
		if unicode.IsSpace(line[idx]) {
			return
		}
	}
	return -1
}
