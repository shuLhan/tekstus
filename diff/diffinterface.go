// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diff

import (
	"bufio"
	//"fmt"
	"github.com/shuLhan/tekstus"
	"io"
	"os"
)

const (
	// LevelLines define that we want only lines change set.
	LevelLines = iota
	// LevelWords define that we want the change not only capture the
	// different per line, but also changes inside the line.
	LevelWords
)

/*
OpenBufferedReader will open the file `f` using buffered reader.
*/
func OpenBufferedReader(f string) (r *bufio.Reader, e error) {
	fd, e := os.Open(f)

	if e == nil {
		r = bufio.NewReader(fd)
	}

	return
}

/*
ReadLines return lines in the file `f`.
*/
func ReadLines(f string) (lines tekstus.Lines, e error) {
	reader, e := OpenBufferedReader(f)

	if e != nil {
		return
	}

	n := 1
	for {
		line, e := reader.ReadBytes(DefDelimiter)

		if e != nil {
			if e == io.EOF {
				break
			}
			return lines, e
		}

		lines = append(lines, tekstus.Line{N: n, V: line})
		n++
	}

	return lines, nil
}

/*
Bytes compare two slice of bytes and return true if equal or false otherwise.
*/
func Bytes(oldb, newb []byte) (equal bool) {
	oldblen := len(oldb)
	newblen := len(newb)

	// Do not compare the length, because we care about the index.

	minlen := 0
	if oldblen < newblen {
		minlen = oldblen
	} else if oldblen == newblen {
		minlen = oldblen
	} else {
		minlen = newblen
	}

	at := 0
	for ; at < minlen; at++ {
		if oldb[at] != newb[at] {
			return
		}
	}

	if oldblen == newblen {
		// Both slice is equal.
		return true
	}

	return false
}

/*
findLine return true if line is found in text beginning at line `startat`.
It also return line number of matching line.
If no match found, it will return false and `startat` value.
*/
func findLine(line tekstus.Line, text tekstus.Lines, startat int) (
	found bool,
	n int,
) {
	textlen := len(text)

	for n = startat; n < textlen; n++ {
		isEqual := Bytes(line.V, text[n].V)

		if isEqual {
			return true, n
		}
	}

	return false, startat
}

/*
Files compare two files.
*/
func Files(oldf, newf string, difflevel int) (diffs Data, e error) {
	oldlines, e := ReadLines(oldf)
	if e != nil {
		return
	}

	newlines, e := ReadLines(newf)
	if e != nil {
		return
	}

	oldlen := len(oldlines)
	newlen := len(newlines)
	x := 0
	y := 0

	for x < oldlen {
		if y == newlen {
			// New text has been full examined. Leave out the old
			// text that means deletion at the end of text.
			diffs.PushDel(oldlines[x])
			oldlines[x].V = nil
			x++
			continue
		}

		// Compare old line with new line.
		isEqual := Bytes(oldlines[x].V, newlines[y].V)

		if isEqual {
			oldlines[x].V = nil
			newlines[y].V = nil
			x++
			y++
			continue
		}

		// x is not equal with y, search down...
		foundx, xaty := findLine(oldlines[x], newlines, y+1)

		// Cross check the y with the rest of x...
		foundy, yatx := findLine(newlines[y], oldlines, x+1)

		// Both line is missing, its mean changes on current line
		if !foundx && !foundy {
			diffs.PushChange(oldlines[x], newlines[y])
			oldlines[x].V = nil
			newlines[y].V = nil
			x++
			y++
			continue
		}

		// x still missing, means deletion in old text.
		if !foundx && foundy {
			for ; x < yatx && x < oldlen; x++ {
				diffs.PushDel(oldlines[x])
				oldlines[x].V = nil
			}
			continue
		}

		// we found x but y is missing, its mean addition in new text.
		if foundx && !foundy {
			//fmt.Printf(">>> foundx at %d while y %d\n", xaty, y)
			for ; y < xaty && y < newlen; y++ {
				diffs.PushAdd(newlines[y])
				newlines[y].V = nil
			}
			continue
		}

		if foundx && foundy {
			// We found x and y. Check which one is the
			// addition or deletion based on line range.
			addlen := xaty - y
			dellen := yatx - x

			if addlen < dellen {
				for ; y < xaty && y < newlen; y++ {
					diffs.PushAdd(newlines[y])
					newlines[y].V = nil
				}
			} else if addlen == dellen {
				// Both changes occur between lines
				for x < yatx && y < xaty {
					diffs.PushChange(oldlines[x],
						newlines[y])
					oldlines[x].V = nil
					newlines[y].V = nil
					x++
					y++
				}
			} else { // addlen > dellen
				for ; x < yatx && x < oldlen; x++ {
					diffs.PushDel(oldlines[x])
					oldlines[x].V = nil
				}
			}
			continue
		}
	}

	// Check if there is a left over from new text.
	for ; y < newlen; y++ {
		diffs.PushAdd(newlines[y])
		newlines[y].V = nil
	}

	return diffs, e
}
