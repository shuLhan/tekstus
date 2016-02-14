// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

import (
	"unicode"
)

/*
StringCountBy count number of occurrence of `class` values in data.
Return number of each class based on their index.

For example, if data is "[A,A,B]" and class is "[A,B]", this function will
return "[2,1]".

	idx cls  count
	0 : A -> 2
	1 : B -> 1
*/
func StringCountBy(data []string, class []string) (clsCnt []int) {
	clsCnt = make([]int, len(class))

	for _, r := range data {
		for k, v := range class {
			if r == v {
				clsCnt[k]++
				break
			}
		}
	}

	return
}

/*
StringsGetMajority return the string that has highest frequency.

Example, given input

	data:  [A A B A B C C]
	class: [A B]

it will return A as the majority class in data.
If class has equal frequency, then the first class in order will returned.
*/
func StringsGetMajority(data []string, class []string) string {
	classCount := StringCountBy(data, class)
	_, maxIdx := IntFindMax(classCount)

	return class[maxIdx]
}

/*
CountCharSequence given a string, count number of repeated character more than
one in sequence and return character and counting value.

Example, given a line of string
	"aaa abcdee ffgf"
it will return
	[a e f]
and
	[3 2 2]

'a' is not counted as 4 because it will breaked by space, so do 'f'.
*/
func CountCharSequence(line string, nospace bool) (chars []rune, counts []int) {
	var lastv rune
	count := 1
	for _, v := range line {
		if v == lastv {
			if !unicode.IsSpace(v) {
				count++
			}
		} else {
			if count > 1 {
				chars = append(chars, lastv)
				counts = append(counts, count)
				count = 1
			}
		}
		lastv = v
	}
	if count > 1 {
		chars = append(chars, lastv)
		counts = append(counts, count)
	}
	return
}

/*
GetMaxCharSequence return character which have maximum sequence in `line`.

Example, given a line of string "aaa abcdee ffgf" it will return 'a' and 3.
*/
func GetMaxCharSequence(line string, nospace bool) (char rune, count int) {
	chars, counts := CountCharSequence(line, nospace)

	if len(chars) == 0 {
		return 0, 0
	}

	_, idx := IntFindMax(counts)

	return chars[idx], counts[idx]
}
