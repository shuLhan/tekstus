// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diff_test

import (
	"github.com/shuLhan/tekstus/diff"
	"testing"
)

type DiffExpect struct {
	Adds    []int
	Dels    []int
	Changes []int
}

type DiffExpects []DiffExpect

func testDiffFiles(t *testing.T, old, new string, level int) diff.Data {
	diffs, e := diff.Files(old, new, level)

	if e != nil {
		t.Fatal(e)
	}

	return diffs
}

func compareLineNumber(t *testing.T, diffs diff.Data, exp DiffExpect) {
	if len(exp.Adds) != len(diffs.Adds) {
		t.Fatalf("Expecting adds at %v, got %v", exp.Adds, diffs.Adds)
	} else {
		for x, v := range exp.Adds {
			if diffs.Adds[x].N != v {
				t.Fatalf("Expecting add at %v, got %v", v,
					diffs.Adds[x])
			}
		}
	}

	if len(exp.Dels) != len(diffs.Dels) {
		t.Fatalf("Expecting deletions at %v, got %v", exp.Dels,
			diffs.Dels)
	} else {
		for x, v := range exp.Dels {
			if diffs.Dels[x].N != v {
				t.Fatalf("Expecting deletion at %v, got %v", v,
					diffs.Dels[x])
			}
		}
	}

	if len(exp.Changes) != len(diffs.Changes) {
		t.Fatalf("Expecting changes at %v, got %v", exp.Changes,
			diffs.Changes)
	} else {
		for x, v := range exp.Changes {
			if diffs.Changes[x].Old.N != v {
				t.Fatalf("Expecting change at %v, got %v", v,
					diffs.Changes[x])
			}
		}
	}
}

func TestDiffFilesLevelLine(t *testing.T) {
	oldrev := "../testdata/328391343.txt"
	newrev := "../testdata/328391582.txt"
	diffsExpects := DiffExpects{
		{[]int{}, []int{}, []int{48}},
		{[]int{}, []int{}, []int{48}},
		{[]int{268, 269, 270, 271}, []int{6, 7, 8, 9, 248, 249, 250},
			[]int{}},
		{[]int{6, 7, 8, 9, 248, 249, 250}, []int{268, 269, 270, 271},
			[]int{}},
		{[]int{54}, []int{},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
				15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
				30, 32, 37, 39, 41, 44, 51},
		},
		{[]int{}, []int{54},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
				15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
				30, 32, 37, 39, 41, 44, 51},
		},
		{[]int{}, []int{5, 6}, []int{}},
		{[]int{5, 6}, []int{}, []int{}},
	}

	diffs := testDiffFiles(t, oldrev, newrev, diff.LevelLines)
	compareLineNumber(t, diffs, diffsExpects[0])

	// reverse test
	diffs = testDiffFiles(t, newrev, oldrev, diff.LevelLines)
	compareLineNumber(t, diffs, diffsExpects[1])

	oldrev = "../testdata/327585467.txt"
	newrev = "../testdata/327607921.txt"

	diffs = testDiffFiles(t, oldrev, newrev, diff.LevelLines)
	compareLineNumber(t, diffs, diffsExpects[2])

	// reverse test
	diffs = testDiffFiles(t, newrev, oldrev, diff.LevelLines)
	compareLineNumber(t, diffs, diffsExpects[3])

	oldrev = "../testdata/314955274.txt"
	newrev = "../testdata/327191082.txt"

	diffs = testDiffFiles(t, oldrev, newrev, diff.LevelLines)
	compareLineNumber(t, diffs, diffsExpects[4])

	diffs = testDiffFiles(t, newrev, oldrev, diff.LevelLines)
	compareLineNumber(t, diffs, diffsExpects[5])

	oldrev = "../testdata/empty5lines.txt"
	newrev = "../testdata/empty3lines.txt"

	diffs = testDiffFiles(t, oldrev, newrev, diff.LevelLines)
	compareLineNumber(t, diffs, diffsExpects[6])

	diffs = testDiffFiles(t, newrev, oldrev, diff.LevelLines)
	compareLineNumber(t, diffs, diffsExpects[7])

}

func TestDiffFilesLevelWords(t *testing.T) {
	t.SkipNow()
	oldrev := "../testdata/328391343.txt"
	newrev := "../testdata/328391582.txt"

	testDiffFiles(t, oldrev, newrev, diff.LevelWords)

	oldrev = "../testdata/327585467.txt"
	newrev = "../testdata/327607921.txt"

	testDiffFiles(t, oldrev, newrev, diff.LevelWords)
}
