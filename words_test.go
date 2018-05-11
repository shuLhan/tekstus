// Copyright 2016-2018 Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"github.com/shuLhan/tekstus"
	"testing"
)

var dataWordsFindLongest = []struct {
	words []string
	exp   string
}{
	{
		[]string{"a", "bb", "ccc", "d", "eee"},
		"ccc",
	}, {
		[]string{"a", "bb", "ccc", "dddd", "eee"},
		"dddd",
	},
}

func TestWordsFindLongest(t *testing.T) {
	for _, td := range dataWordsFindLongest {
		got, _ := tekstus.WordsFindLongest(td.words)

		assert(t, td.exp, got, true)
	}
}
