// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

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
