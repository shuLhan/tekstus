// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package diff implement text comparison.
*/
package diff

import (
	"fmt"
	"github.com/shuLhan/tekstus"
)

var (
	// DefDelimiter define default delimiter for new line.
	DefDelimiter = byte('\n')
)

/*
Data represent additions, deletions, and changes between two text.
*/
type Data struct {
	Adds    tekstus.Lines
	Dels    tekstus.Lines
	Changes LineChanges
}

/*
PushAdd will add new line to diff set.
*/
func (diffs *Data) PushAdd(new tekstus.Line) {
	diffs.Adds = append(diffs.Adds, new)
}

/*
PushDel will add deletion line to diff set.
*/
func (diffs *Data) PushDel(old tekstus.Line) {
	diffs.Dels = append(diffs.Dels, old)
}

/*
PushChange set to diff data.
*/
func (diffs *Data) PushChange(old, new tekstus.Line) {
	change := NewLineChange(old, new)

	diffs.Changes = append(diffs.Changes, *change)
}

/*
GetAllAdds return chunks of additions including in line changes.
*/
func (diffs *Data) GetAllAdds() (chunks tekstus.Chunks) {
	for _, add := range diffs.Adds {
		chunks = append(chunks, tekstus.Chunk{0, add.V})
	}
	chunks = append(chunks, diffs.Changes.GetAllAdds()...)
	return
}

/*
GetAllDels return chunks of deletions including in line changes.
*/
func (diffs *Data) GetAllDels() (chunks tekstus.Chunks) {
	for _, del := range diffs.Dels {
		chunks = append(chunks, tekstus.Chunk{0, del.V})
	}
	chunks = append(chunks, diffs.Changes.GetAllDels()...)
	return
}

/*
String return formatted data.
*/
func (diffs Data) String() (s string) {
	s += "Diffs:\n"

	if len(diffs.Adds) > 0 {
		s += ">>> Adds:\n"
		for _, add := range diffs.Adds {
			s += fmt.Sprintf("  + %d : %s", add.N, string(add.V))
		}
	}

	if len(diffs.Dels) > 0 {
		s += ">>> Dels:\n"
		for _, del := range diffs.Dels {
			s += fmt.Sprintf("  - %d : %s", del.N, string(del.V))
		}
	}

	if len(diffs.Changes) > 0 {
		s += ">>> Changes:\n" + fmt.Sprint(diffs.Changes)
	}

	return
}
