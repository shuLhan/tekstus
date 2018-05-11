// Copyright 2016-2018 Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diff

import (
	"github.com/shuLhan/tekstus"
)

//
// LineChanges represet a set of change in text.
//
type LineChanges []LineChange

//
// GetAllDels return all deleted chunks.
//
func (changes *LineChanges) GetAllDels() (allDels tekstus.Chunks) {
	for _, change := range *changes {
		allDels = append(allDels, change.Dels...)
	}
	return
}

//
// GetAllAdds return all addition chunks.
//
func (changes *LineChanges) GetAllAdds() (allAdds tekstus.Chunks) {
	for _, change := range *changes {
		allAdds = append(allAdds, change.Adds...)
	}
	return
}
