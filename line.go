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
String return formatted data.
*/
func (line Line) String() string {
	return fmt.Sprintf("Line: { N: %d, V: %s }\n", line.N, string(line.V))
}
