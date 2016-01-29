// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"github.com/shuLhan/tekstus"
	"testing"
)

var tssExp = tekstus.TableStrings{
	{{"a"}, {"b", "c"}},
	{{"b"}, {"a", "c"}},
	{{"c"}, {"a", "b"}},
}

func TestTableStringsIsEqual(t *testing.T) {
	b := tssExp.IsEqual(tssExp)
	if !b {
		t.Fatal("Expecting true, got", tssExp, " == ", tssExp, "? ", b)
	}

	setstr := tekstus.TableStrings{
		{{"c"}, {"a", "b"}},
		{{"a"}, {"b", "c"}},
		{{"b"}, {"a", "c"}},
	}

	b = tssExp.IsEqual(setstr)
	if !b {
		t.Fatal("Expecting true, got", tssExp, " == ", setstr, "? ", b)
	}

	setstr = tekstus.TableStrings{
		{{"c"}, {"a", "b"}},
		{{"a"}, {"b", "c"}},
	}

	b = tssExp.IsEqual(setstr)
	if b {
		t.Fatal("Expecting false, got", tssExp, " == ", setstr, "? ", b)
	}

	setstr = tekstus.TableStrings{
		{{"b"}, {"a", "b"}},
		{{"c"}, {"a", "b"}},
		{{"a"}, {"b", "c"}},
	}

	b = tssExp.IsEqual(setstr)
	if b {
		t.Fatal("Expecting false, got", tssExp, " == ", setstr, "? ", b)
	}
}
