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

/*
CountUpperLowerChar return number of uppercase and lowercase in line.
*/
func CountUpperLowerChar(line string) (upper, lower int) {
	for _, v := range line {
		if !unicode.IsLetter(v) {
			continue
		}
		if unicode.IsUpper(v) {
			upper++
		} else {
			lower++
		}
	}
	return
}

/*
RatioUpperLowerChar compute and return ratio of uppercase with lowercase
character in line.
*/
func RatioUpperLowerChar(line string) float32 {
	up, lo := CountUpperLowerChar(line)

	return float32(1+up) / float32(1+lo)
}

/*
RatioUpper compute and return ratio of uppercase character to all character in
line.
*/
func RatioUpper(line string) float32 {
	up, lo := CountUpperLowerChar(line)

	return float32(1+up) / float32(1+up+lo)
}

/*
CountDigit return number of digit in line.
*/
func CountDigit(line string) (n int) {
	for _, v := range line {
		if unicode.IsDigit(v) {
			n++
		}
	}
	return
}

/*
RatioDigit compute and return digit ratio to all characters in line.
*/
func RatioDigit(line string) float32 {
	n := CountDigit(line)
	slen := len(line)

	return float32(1+n) / float32(1+slen)
}

/*
CountAlnumChar return number of alpha-numeric character in line and length of
line.
*/
func CountAlnumChar(line string) (n, l int) {
	l = len(line)
	for _, v := range line {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			n++
		}
	}
	return
}

/*
RatioAlnumChar compute and return ratio of alpha-numeric with all character in
line.
*/
func RatioAlnumChar(line string) float32 {
	n, length := CountAlnumChar(line)

	return float32(1+n) / float32(1+length)
}

/*
CountUniqChar count number of character in line without duplication and return
it along with length of line.

Example, if line is "aba" then it will count as 2 ("a", "b").
*/
func CountUniqChar(line string) (n, l int) {
	var uchars []rune

	l = len(line)

	for _, v := range line {
		yes, _ := RunesContain(uchars, v)
		if yes {
			continue
		}
		uchars = append(uchars, v)
		n++
	}
	return
}

/*
CountAlnumDistribution count alpha-numeric characters in line.

Example, given a line "abbcccddddeeeee", it will return [a b c d e] and
[1 2 3 4 5].
*/
func CountAlnumDistribution(line string) (chars []rune, values []int) {
	var found = false

	for _, v := range line {
		if !(unicode.IsDigit(v) || unicode.IsLetter(v)) {
			continue
		}
		found = false
		for y, c := range chars {
			if v == c {
				values[y]++
				found = true
				break
			}
		}
		if !found {
			chars = append(chars, v)
			values = append(values, 1)
		}
	}
	return
}
