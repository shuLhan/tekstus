// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

/*
FindToken return the first index of matched token in line.
If not found it will return -1.
*/
func FindToken(token, line []byte, startat int) (at int) {
	y := 0
	tokenlen := len(token)
	linelen := len(line)

	at = -1
	for x := startat; x < linelen; x++ {
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
EncapsulateToken will find `token` in `line` and capsulating it with bytes
from `leftcap` and `rightcap`.
If no token is found, it will return the same line with false status.
*/
func EncapsulateToken(token, line, leftcap, rightcap []byte) (
	newline []byte,
	changed bool,
) {
	tokenlen := len(token)

	startat := 0
	for {
		foundat := FindToken(token, line, startat)

		if foundat < 0 {
			newline = append(newline, line[startat:]...)
			break
		}

		newline = append(newline, line[startat:foundat]...)
		newline = append(newline, leftcap...)
		newline = append(newline, token...)
		newline = append(newline, rightcap...)

		startat = foundat + tokenlen
	}

	if startat > 0 {
		changed = true
	}

	return
}
