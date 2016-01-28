// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

import (
	"fmt"
)

/*
Line represent bytes of string and line number.
*/
type Line struct {
	N int
	V []byte
}

/*
Lines represent array of line.
*/
type Lines []Line

/*
NewLine create and return pointer to new Line object.
*/
func NewLine(n int, v []byte) *Line {
	return &Line{n, v}
}

/*
FindToken return the first index of matched token in line.
If not found it will return -1.
*/
func FindToken(token, line []byte) (at int) {
	y := 0
	tokenlen := len(token)
	linelen := len(line)

	at = -1
	for x := 0; x < linelen; x++ {
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
			// reset back
			y = 0
			at = -1
		}
	}
	// x run out before y
	if y < tokenlen {
		at = -1
	}
	return
}

/*
String return formatted data.
*/
func (line Line) String() string {
	return fmt.Sprintf("Line: { N: %d, V: %s }\n", line.N, string(line.V))
}
