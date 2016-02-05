// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diff

import (
	"fmt"
	"github.com/shuLhan/tekstus"
)

/*
LineChange represent one change in text.
*/
type LineChange struct {
	Old  tekstus.Line
	New  tekstus.Line
	Adds tekstus.Chunks
	Dels tekstus.Chunks
}

/*
NewLineChange create a pointer to new LineChange object.
*/
func NewLineChange(old, new tekstus.Line) *LineChange {
	return &LineChange{old, new, tekstus.Chunks{}, tekstus.Chunks{}}
}

/*
String return formatted content of LineChange.
*/
func (change LineChange) String() string {
	return fmt.Sprintf("LineChange: {\n"+
		" Old  : %v\n"+
		" New  : %v\n"+
		" Adds : %v\n"+
		" Dels : %v\n"+
		"}\n", change.Old, change.New, change.Adds, change.Dels)
}
