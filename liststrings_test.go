// Copyright 2016-2018 Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"github.com/shuLhan/tekstus"
	"testing"
)

func TestListStringsIsEqual(t *testing.T) {
	b := tssExp[0].IsEqual(tssExp[0])
	if !b {
		t.Fatalf("Expecting true, got %v == %v ? %v", tssExp[0],
			tssExp[0], b)
	}

	subset := tekstus.ListStrings{{"a"}, {"c", "b"}}
	b = tssExp[0].IsEqual(subset)
	if !b {
		t.Fatal("Expecting true, got", tssExp[0], " == ", subset, "? ", b)
	}

	subset = tekstus.ListStrings{{"a"}, {"b", "a"}}
	b = tssExp[0].IsEqual(subset)
	if b {
		t.Fatal("Expecting false, got", tssExp[0], " == ", subset, "? ", b)
	}

	b = tssExp[0].IsEqual(tssExp[1])
	if b {
		t.Fatal("Expecting false, got", tssExp[0], " == ", tssExp[1], "? ", b)
	}
}
