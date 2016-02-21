// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"reflect"
	"runtime/debug"
	"testing"
)

func assert(t *testing.T, exp, got interface{}, equal bool) {
	if reflect.DeepEqual(exp, got) != equal {
		debug.PrintStack()
		t.Fatalf("\n"+
			">>> Expecting '%v'\n"+
			"          got '%v'\n", exp, got)
	}
}

var dataLines = []string{
	"// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.",
	"// Copyright 2016 Mhd Sulhan \"<ms@kilabit.info>\". All rights reserved.",
	"/* TEST */",
	"ftp://test.com/123 The [[United States]] has regularly voted alone and against international consensus, using its [[United Nations Security Council veto power|veto power]] to block the adoption of proposed UN Security Council resolutions supporting the [[PLO]] and calling for a two-state solution to the [[Israeli-Palestinian conflict]].<ref>[http://books.google.ca/books?id=CHL5SwGvobQC&pg=PA168&dq=US+veto+Israel+regularly#v=onepage&q=US%20veto%20Israel%20regularly&f=false Pirates and emperors, old and new: international terrorism in the real world], [[Noam Chomsky]], p. 168.</ref><ref>The US has also used its veto to block resolutions that are critical of Israel.[https://books.google.ca/books?id=yzmpDAz7ZAwC&pg=PT251&dq=US+veto+Israel+regularly&lr=#v=onepage&q=US%20veto%20Israel%20regularly&f=false Uneasy neighbors], David T. Jones and David Kilgour, p. 235.</ref> The United States responded to the frequent criticism from UN organs by adopting the [[Negroponte doctrine]].",
	"The [[United States]] has regularly voted alone and against international consensus, using its [[United Nations Security Council veto power|veto power]] to block the adoption of proposed UN Security Council resolutions supporting the [[PLO]] and calling for a two-state solution to the [[Israeli-Palestinian conflict]].",
	"   a  b c   d  ",
	"   a\n\nb c   d\n\n",
	`==External links==
*[http://www.bigfinish.com/24-Doctor-Who-The-Eye-of-the-Scorpion Big Finish Productions - ''The Eye of the Scorpion'']
*{{Doctor Who RG | id=who_bf24 | title=The Eye of the Scorpion}}
===Reviews===
* Test image [[Image:fileto.png]].
* Test file [[File:fileto.png]].
*{{OG review | id=bf-24 | title=The Eye of the Scorpion}}
*{{DWRG | id=eyes | title=The Eye of the Scorpion}}

<br clear="all">
{{Fifthdoctoraudios}}

{{DEFAULTSORT:Eye of the Scoprion, The}}
[[Category:Fifth Doctor audio plays]]
[[:Category:2001 audio plays]]
{{DoctorWho-stub}}`,
}

var dataEncapsulateTrim = []struct {
	text     string
	leftcap  string
	rightcap string
	exp      string
}{
	{
		dataLines[1], "<", ">",
		"// Copyright 2016 Mhd Sulhan \"\". All rights reserved.",
	}, {
		dataLines[1], "\"", "\"",
		"// Copyright 2016 Mhd Sulhan . All rights reserved.",
	}, {
		dataLines[1], "//", "//",
		"// Copyright 2016 Mhd Sulhan \"<ms@kilabit.info>\". All rights reserved.",
	}, {
		dataLines[2], "/*", "*/", "",
	},
}
