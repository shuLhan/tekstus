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
	Old tekstus.Line
	New tekstus.Line
}

/*
LineChanges represet a set of change in text.
*/
type LineChanges []LineChange

/*
NewLineChange create a pointer to new LineChange object.
*/
func NewLineChange(old, new tekstus.Line) *LineChange {
	return &LineChange{old, new}
}

/*
String return formatted content of LineChange.
*/
func (change LineChange) String() string {
	return fmt.Sprintf("LineChange {\n"+
		" old  : %v\n"+
		" new  : %v\n"+
		"}", change.Old, change.New)
}
