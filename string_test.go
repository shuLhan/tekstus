// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"github.com/shuLhan/tekstus"
	"testing"
)

var dataStringCountTokens = []struct {
	line   string
	tokens []string
	exp    int
}{
	{
		dataLines[0],
		[]string{"//"},
		1,
	}, {
		dataLines[3],
		[]string{"[[", "]]", "<ref", "/ref>", "[http:"},
		17,
	},
}

func TestStringCountTokens(t *testing.T) {
	for _, td := range dataStringCountTokens {
		got := tekstus.StringCountTokens(td.line, td.tokens, false)

		assert(t, td.exp, got, true)
	}
}

var dataStringTrimNonAlnum = []struct {
	text string
	exp  string
}{
	{"[[alpha]]", "alpha"},
	{"[[alpha", "alpha"},
	{"alpha]]", "alpha"},
	{"alpha", "alpha"},
	{"alpha0", "alpha0"},
	{"1alpha", "1alpha"},
	{"1alpha0", "1alpha0"},
	{"[][][]", ""},
}

func TestStringTrimNonAlnum(t *testing.T) {
	for _, td := range dataStringTrimNonAlnum {
		got := tekstus.StringTrimNonAlnum(td.text)

		assert(t, td.exp, got, true)
	}
}

var dataStringSplitWords = []struct {
	text string
	exp  []string
}{
	{
		dataLines[0],
		[]string{"Copyright", "2016", "Mhd", "Sulhan",
			"ms@kilabit.info", "All", "rights", "reserved"},
	},
	{
		dataLines[4],
		[]string{"The", "United", "States", "has", "regularly",
			"voted", "alone", "and", "against", "international",
			"consensus", "using", "its", "Nations", "Security",
			"Council", "veto", "power|veto", "power", "to",
			"block", "adoption", "of", "proposed", "UN",
			"resolutions", "supporting", "PLO", "calling",
			"for", "a", "two-state", "solution",
			"Israeli-Palestinian", "conflict",
		},
	},
}

func TestStringSplitWords(t *testing.T) {
	for _, td := range dataStringSplitWords {
		got := tekstus.StringSplitWords(td.text, true, true)

		assert(t, td.exp, got, true)
	}
}

var dataStringRemoveURI = []struct {
	text string
	exp  string
}{
	{
		dataLines[3],
		" The [[United States]] has regularly voted alone and against international consensus, using its [[United Nations Security Council veto power|veto power]] to block the adoption of proposed UN Security Council resolutions supporting the [[PLO]] and calling for a two-state solution to the [[Israeli-Palestinian conflict]].<ref>[ Pirates and emperors, old and new: international terrorism in the real world], [[Noam Chomsky]], p. 168.</ref><ref>The US has also used its veto to block resolutions that are critical of Israel.[ Uneasy neighbors], David T. Jones and David Kilgour, p. 235.</ref> The United States responded to the frequent criticism from UN organs by adopting the [[Negroponte doctrine]].",
	},
}

func TestStringRemoveURI(t *testing.T) {
	for _, td := range dataStringRemoveURI {
		got := tekstus.StringRemoveURI(td.text)

		assert(t, td.exp, got, true)
	}
}

var dataStringMergeSpaces = []struct {
	text string
	exp  string
}{
	{
		dataLines[5],
		" a b c d ",
	},
	{
		dataLines[6],
		" a\nb c d\n",
	},
}

func TestStringMergeSpaces(t *testing.T) {
	for _, td := range dataStringMergeSpaces {
		got := tekstus.StringMergeSpaces(td.text, true)

		assert(t, td.exp, got, true)
	}
}

var dataStringRemoveWikiMarkup = []struct {
	text string
	exp  string
}{
	{
		dataLines[7],
		`==External links==
*[http://www.bigfinish.com/24-Doctor-Who-The-Eye-of-the-Scorpion Big Finish Productions - ''The Eye of the Scorpion'']
*{{Doctor Who RG | id=who_bf24 | title=The Eye of the Scorpion}}
===Reviews===
* Test image .
* Test file .
*{{OG review | id=bf-24 | title=The Eye of the Scorpion}}
*{{DWRG | id=eyes | title=The Eye of the Scorpion}}

<br clear="all">
{{Fifthdoctoraudios}}




{{DoctorWho-stub}}`,
	},
}

func TestStringRemoveWikiMarkup(t *testing.T) {
	for _, td := range dataStringRemoveWikiMarkup {
		got := tekstus.StringRemoveWikiMarkup(td.text)

		assert(t, td.exp, got, true)
	}
}
