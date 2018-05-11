// Copyright 2016-2018 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
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
RunesDiff return the difference between two slice of rune.

For example, input are

	l: [a b c d]
	r: [b c]

and the output will be `[a d]`
*/
func RunesDiff(l []rune, r []rune) (diff []rune) {
	var found bool
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
RunesFind will search token in text starting from index `startAt` and
return the matching index.

If no token is found it will return -1.
*/
func RunesFind(line, token []rune, startAt int) (at int) {
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

/*
RunesRemoveUntil given a line, remove all characters inside it, starting
from `leftcap` until the `rightcap` and return cutted line and changed to true.

If no `leftcap` or `rightcap` is found, the line will unchanged, and changed
will be false.

Example,

	line    : "[[ ABC ]] DEF"
	leftcap : "[["
	rightcap: "]]"
	return  : "  DEF"
*/
func RunesRemoveUntil(line, leftcap, rightcap []rune) (
	newline []rune,
	changed bool,
) {
	lidx := RunesFind(line, leftcap, 0)
	ridx := RunesFind(line, rightcap, lidx+1)

	if lidx < 0 || ridx < 0 || lidx >= ridx {
		return line, false
	}

	newline = line[:lidx]
	newline = append(newline, line[ridx+len(rightcap):]...)
	newline, _ = RunesRemoveUntil(newline, leftcap, rightcap)

	return newline, true
}
