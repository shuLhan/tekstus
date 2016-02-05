// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

import (
	"fmt"
)

/*
Chunk represent subset of line, contain starting position and slice of bytes in
line.
*/
type Chunk struct {
	StartAt int
	V       []byte
}

/*
String return formatted data.
*/
func (chunk Chunk) String() string {
	return fmt.Sprintf("{ %d {%s}}", chunk.StartAt, string(chunk.V))
}
