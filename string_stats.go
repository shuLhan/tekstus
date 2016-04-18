// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

import (
	"github.com/shuLhan/numerus"
	"strings"
	"unicode"
)

/*
CountCharSequence given a string, count number of repeated character more than
one in sequence and return character and counting value.

Example, given a text of string
	"aaa abcdee ffgf"
it will return
	[a e f]
and
	[3 2 2]

'a' is not counted as 4 because it will breaked by space, so do 'f'.
*/
func CountCharSequence(text string) (chars []rune, counts []int) {
	var lastv rune
	count := 1
	for _, v := range text {
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
GetMaxCharSequence return character which have maximum sequence in `text`.

Example, given a text of string "aaa abcdee ffgf" it will return 'a' and 3.
*/
func GetMaxCharSequence(text string) (char rune, count int) {
	chars, counts := CountCharSequence(text)

	if len(chars) == 0 {
		return 0, 0
	}

	_, idx, _ := numerus.IntsFindMax(counts)

	return chars[idx], counts[idx]
}

/*
CountUpperLowerChar return number of uppercase and lowercase in text.
*/
func CountUpperLowerChar(text string) (upper, lower int) {
	for _, v := range text {
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
character in text.
*/
func RatioUpperLowerChar(text string) float64 {
	if len(text) == 0 {
		return 0
	}

	up, lo := CountUpperLowerChar(text)

	if lo == 0 {
		return float64(up)
	}

	return float64(up) / float64(lo)
}

/*
RatioUpper compute and return ratio of uppercase character to all character in
text.
*/
func RatioUpper(text string) float64 {
	if len(text) == 0 {
		return 0
	}
	up, lo := CountUpperLowerChar(text)

	total := up + lo
	if total == 0 {
		return 0
	}

	return float64(up) / float64(total)
}

/*
CountDigit return number of digit in text.
*/
func CountDigit(text string) (n int) {
	if len(text) == 0 {
		return 0
	}

	for _, v := range text {
		if unicode.IsDigit(v) {
			n++
		}
	}
	return
}

/*
RatioDigit compute and return digit ratio to all characters in text.
*/
func RatioDigit(text string) float64 {
	textlen := len(text)

	if textlen == 0 {
		return 0
	}

	n := CountDigit(text)

	if n == 0 {
		return 0
	}

	return float64(n) / float64(textlen)
}

/*
CountAlnumChar return number of alpha-numeric character in text.
*/
func CountAlnumChar(text string) (n int) {
	if len(text) == 0 {
		return
	}

	for _, v := range text {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			n++
		}
	}
	return
}

/*
RatioAlnumChar compute and return ratio of alpha-numeric with all character in
text.
*/
func RatioAlnumChar(text string) float64 {
	textlen := len(text)
	if textlen == 0 {
		return 0
	}

	n := CountAlnumChar(text)

	return float64(n) / float64(textlen)
}

/*
CountNonAlnumChar return number of non alpha-numeric character in text.
If `withspace` is true, it will be counted as non-alpha-numeric, if it false
it will be skipped.
*/
func CountNonAlnumChar(text string, withspace bool) (n int) {
	if len(text) == 0 {
		return
	}

	for _, v := range text {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			continue
		}
		if unicode.IsSpace(v) {
			if withspace {
				n++
			}
			continue
		}
		n++
	}
	return
}

/*
RatioNonAlnumChar return ratio of non-alphanumeric character to all character
in text.
If `withspace` is true then white-space character will be counted as non-alpha
numeric, otherwise it will be skipped.
*/
func RatioNonAlnumChar(text string, withspace bool) float64 {
	textlen := len(text)
	if textlen == 0 {
		return 0
	}

	n := CountNonAlnumChar(text, withspace)

	return float64(n) / float64(textlen)
}

/*
CountUniqChar count number of character in text without duplication.

Example, if text is "aba" then it will count as 2 ("a", "b").
*/
func CountUniqChar(text string) (n int) {
	textlen := len(text)
	if textlen == 0 {
		return
	}

	var uchars []rune

	for _, v := range text {
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
CountAlnumDistribution count distribution of alpha-numeric characters in text.

Example, given a text "abbcccddddeeeee", it will return [a b c d e] and
[1 2 3 4 5].
*/
func CountAlnumDistribution(text string) (chars []rune, values []int) {
	var found = false

	for _, v := range text {
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

/*
StringCountTokens given a text, count how many tokens inside of it and return
sum of all.
*/
func StringCountTokens(text string, tokens []string, sensitive bool) (
	cnt int,
) {
	if len(text) == 0 {
		return 0
	}

	if !sensitive {
		text = strings.ToLower(text)
	}

	for _, v := range tokens {
		if !sensitive {
			v = strings.ToLower(v)
		}
		cnt += strings.Count(text, v)
	}

	return
}

/*
StringFrequenciesOf return frequencies of tokens by counting each occurence
of token and divide it with total words in text.
*/
func StringFrequenciesOf(text string, tokens []string, sensitive bool) (
	freq float64,
) {
	if len(text) <= 0 {
		return 0
	}

	textWords := StringSplitWords(text, false, false)
	textWordsLen := float64(len(textWords))

	tokensCnt := float64(StringCountTokens(text, tokens, sensitive))

	freq = tokensCnt / textWordsLen

	return
}
