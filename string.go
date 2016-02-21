// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

import (
	"strings"
	"unicode"
)

/*
StringCountTokens given a text, count how many tokens inside of it and return
sum of all.
*/
func StringCountTokens(text string, tokens []string) (cnt int) {
	for _, v := range tokens {
		cnt += strings.Count(text, v)
	}
	return
}

/*
StringTrimNonAlnum remove non alpha-numeric character at the beginning and end
for `text`.
*/
func StringTrimNonAlnum(text string) string {
	r := []rune(text)
	rlen := len(r)
	start := 0

	for ; start < rlen; start++ {
		if unicode.IsLetter(r[start]) || unicode.IsDigit(r[start]) {
			break
		}
	}

	if start >= rlen {
		return ""
	}

	r = r[start:]
	rlen = len(r)
	end := rlen - 1
	for ; end >= 0; end-- {
		if unicode.IsLetter(r[end]) || unicode.IsDigit(r[end]) {
			break
		}
	}

	if end < 0 {
		return ""
	}

	r = r[:end+1]

	return string(r)
}

/*
StringUniq remove duplicate word from `words`. If sensitive is true then
compare the string with case sensitive.

Return the list of unique words.
*/
func StringUniq(words []string, sensitive bool) (uniques []string) {
	// Remove duplicate values.
	wordslen := len(words)
	xcmp := ""
	ycmp := ""

	for x, v := range words {
		if v == "" {
			continue
		}

		if sensitive {
			xcmp = v
		} else {
			xcmp = strings.ToLower(v)
		}

		for y := x + 1; y < wordslen; y++ {
			if len(words[y]) == 0 {
				continue
			}

			if sensitive {
				ycmp = words[y]
			} else {
				ycmp = strings.ToLower(words[y])
			}

			if xcmp == ycmp {
				words[y] = ""
			}
		}

		uniques = append(uniques, v)
	}
	return
}

/*
StringSplitWords given a text, return all words in text.

Definition of word are,
- any sequence of characters that is equal or greater than one that is
separated by space.
- does not start with number
- does not end with number

If cleanit is true remove any non-alphanumeric in the start and the end of
each words.

If uniq is true remove duplicate words.
*/
func StringSplitWords(text string, cleanit bool, uniq bool) (words []string) {
	words = strings.Fields(text)

	if !cleanit {
		return
	}

	// Clean the fields, remove non-alphanumeric character from start and
	// end.
	for x, word := range words {
		words[x] = StringTrimNonAlnum(word)
	}

	if !uniq {
		return
	}

	// Remove duplicate values.
	return StringUniq(words, false)
}

/*
StringRemoveURI remove link (http, https, ftp, ftps) from text and return the
new text.
This function assume that space in URI is using '%20'.
*/
func StringRemoveURI(text string) string {
	if len(text) <= 0 {
		return ""
	}

	var uris = []string{
		"http://", "https://", "ftp://", "ftps://",
	}

	ctext := []rune(text)

	for _, uri := range uris {
		startat := 0
		curi := []rune(uri)
		newtext := []rune{}

		for {
			begin := RunesFindToken(ctext, curi, startat)
			if begin < 0 {
				if startat > 0 {
					newtext = append(newtext,
						ctext[startat:]...)
				}
				break
			}

			newtext = append(newtext, ctext[startat:begin]...)

			end := RunesFindSpaces(ctext, begin)
			if end < 0 {
				break
			}

			startat = end
		}
		if len(newtext) > 0 {
			ctext = newtext
		}
	}
	return string(ctext)
}

/*
StringMergeSpaces replace two or more spaces with single space. If withline is
true it also replace two or more new lines with single new-line.
*/
func StringMergeSpaces(text string, withline bool) string {
	var out []rune
	var isspace bool
	var isnewline bool

	for _, v := range text {
		if v == ' ' {
			if isspace {
				continue
			}
			isspace = true
		} else {
			if isspace {
				isspace = false
			}
		}
		if withline {
			if v == '\n' {
				if isnewline {
					continue
				}
				isnewline = true
			} else {
				if isnewline {
					isnewline = false
				}
			}
		}
		out = append(out, v)
	}
	return string(out)
}

/*
StringRemoveWikiMarkup remove wiki markup, including,
- [[Category: ... ]]
- [[:Category: ... ]]
- [[File: ... ]]
- [[Help: ... ]]
- [[Image: ... ]]
- [[Special: ... ]]
- [[Wikipedia: ... ]]
- {{DEFAULTSORT: ... }}
- {{Template: ... }}
- <ref ... />
*/
func StringRemoveWikiMarkup(text string) string {
	markups := []struct {
		begin []rune
		end   []rune
	}{
		{
			[]rune("[[Category:"),
			[]rune("]]"),
		}, {
			[]rune("[[:Category:"),
			[]rune("]]"),
		}, {
			[]rune("[[File:"),
			[]rune("]]"),
		}, {
			[]rune("[[Help:"),
			[]rune("]]"),
		}, {
			[]rune("[[Image:"),
			[]rune("]]"),
		}, {
			[]rune("[[Special:"),
			[]rune("]]"),
		}, {
			[]rune("[[Wikipedia:"),
			[]rune("]]"),
		}, {
			[]rune("{{DEFAULTSORT:"),
			[]rune("}}"),
		}, {
			[]rune("{{Template:"),
			[]rune("}}"),
		}, {
			[]rune("<ref"),
			[]rune("/>"),
		},
	}

	ctext := []rune(text)

	for _, mu := range markups {
		ctext, _ = RunesEncapsulateTrim(ctext, mu.begin, mu.end)
	}

	return string(ctext)
}
