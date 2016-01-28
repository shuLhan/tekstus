// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diff_test

import (
	"github.com/shuLhan/tekstus"
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

	oldrev := "../testdata/Top_Gear_Series_14.old"
	newrev := "../testdata/Top_Gear_Series_14.new"

	diffs := testDiffFiles(t, oldrev, newrev, diff.LevelLines)
	compareLineNumber(t, diffs, diffsExpects[0])

	// reverse test
	diffs = testDiffFiles(t, newrev, oldrev, diff.LevelLines)
	compareLineNumber(t, diffs, diffsExpects[1])

	oldrev = "../testdata/List_of_United_Nations.old"
	newrev = "../testdata/List_of_United_Nations.new"

	diffs = testDiffFiles(t, oldrev, newrev, diff.LevelLines)
	compareLineNumber(t, diffs, diffsExpects[2])

	// reverse test
	diffs = testDiffFiles(t, newrev, oldrev, diff.LevelLines)
	compareLineNumber(t, diffs, diffsExpects[3])

	oldrev = "../testdata/Psusennes_II.old"
	newrev = "../testdata/Psusennes_II.new"

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
	exp_adds := [][]string{
		[]string{"pharaoh"},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"|"},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"|"},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{"| "},
		[]string{" name=\"Kitchen, p.423\""},
		[]string{" name=\"Payraudeau, BIFAO 108, p.294\"", "—", "—",
			" name=\"", "\"/",
		},
		[]string{" name=\"Kitchen, p.290\"", " name=\"", "\"/", "–",
			"—", "—",
		},
		[]string{"—"},
		[]string{
			"—",
			" name=\"Krauss, DE 62, pp.43-48\"",
			" name=\"",
			"\"/",
		},
		[]string{"—", "—", "—", " name=\"", "\"/", "—"},
		[]string{"&nbsp;"},
	}

	exp_dels := [][]string{
		[]string{"Pharaoh ", "| "},
		[]string{"   ", " ", " |"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "  |"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", " |"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", "|"},
		[]string{"   ", " ", " |"},
		[]string{},
		[]string{"--", "--", ">", "</ref"},
		[]string{">", "</ref", "-", "--", "--"},
		[]string{"--"},
		[]string{"--", ">", "</ref"},
		[]string{"--", "--", "--", ">", "</ref", "--"},
		[]string{},
	}

	oldrev := "../testdata/text01.old"
	newrev := "../testdata/text01.new"

	diffs := testDiffFiles(t, oldrev, newrev, diff.LevelWords)

	compareChunks(t, diffs.Changes[0].Adds, diffs.Changes[0].Dels,
		exp_adds[26], exp_dels[26])

	oldrev = "../testdata/text02.old"
	newrev = "../testdata/text02.new"

	diffs = testDiffFiles(t, oldrev, newrev, diff.LevelWords)
	compareChunks(t, diffs.Changes[0].Adds, diffs.Changes[0].Dels,
		exp_adds[27], exp_dels[27])

	oldrev = "../testdata/Top_Gear_Series_14.old"
	newrev = "../testdata/Top_Gear_Series_14.new"

	diffs = testDiffFiles(t, oldrev, newrev, diff.LevelWords)
	compareChunks(t, diffs.Changes[0].Adds, diffs.Changes[0].Dels,
		[]string{","},
		[]string{"alse "},
	)

	oldrev = "../testdata/Psusennes_II.old"
	newrev = "../testdata/Psusennes_II.new"

	diffs = testDiffFiles(t, oldrev, newrev, diff.LevelWords)
	for x, change := range diffs.Changes {
		if x >= len(exp_adds) {
			break
		}
		compareChunks(t, change.Adds, change.Dels, exp_adds[x],
			exp_dels[x])
	}
}

func compareChunks(t *testing.T, adds, dels tekstus.Chunks,
	exp_adds, exp_dels []string,
) {
	if len(adds) != len(exp_adds) {
		t.Fatalf("Expecting adds '%v' got '%v'", exp_adds, adds)
	}
	for x, add := range adds {
		addv := string(add.V)
		if addv != exp_adds[x] {
			t.Fatalf("[%d] Expecting add '%v' got '%v'", x,
				exp_adds[x], addv)
		}
	}

	if len(dels) != len(exp_dels) {
		t.Fatalf("Expecting deletes '%v' got '%v'", exp_dels, dels)
	}
	for x, del := range dels {
		delv := string(del.V)
		if delv != exp_dels[x] {
			t.Fatalf("[%d] Expecting delete '%v' got '%v'", x,
				exp_dels[x], delv)
		}
	}
}

func testDiffLines(t *testing.T, old, new tekstus.Line,
	exp_adds, exp_dels []string) {

	adds, dels := diff.Lines(old.V, new.V, 0, 0)

	compareChunks(t, adds, dels, exp_adds, exp_dels)
}

func TestDiffLines(t *testing.T) {
	old := tekstus.Line{N: 0, V: []byte("lorem ipsum dolmet")}
	new := tekstus.Line{N: 0, V: []byte("lorem all ipsum")}

	exp_adds := [][]string{
		[]string{"all "},
	}
	exp_dels := [][]string{
		[]string{" dolmet"},
	}

	testDiffLines(t, old, new, exp_adds[0], exp_dels[0])

	old = tekstus.Line{N: 0, V: []byte("lorem ipsum dolmet")}
	new = tekstus.Line{N: 0, V: []byte("lorem ipsum")}

	testDiffLines(t, old, new, []string{}, exp_dels[0])

	old = tekstus.Line{N: 0, V: []byte("lorem ipsum")}
	new = tekstus.Line{N: 0, V: []byte("lorem ipsum dolmet")}

	testDiffLines(t, old, new, exp_dels[0], []string{})

	old = tekstus.Line{N: 0, V: []byte("{{Pharaoh Infobox |")}
	new = tekstus.Line{N: 0, V: []byte("{{Infobox pharaoh")}

	testDiffLines(t, old, new, []string{"pharaoh"},
		[]string{"Pharaoh ", "|"})
}
