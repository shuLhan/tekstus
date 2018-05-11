// Copyright 2016-2018 Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"github.com/shuLhan/tekstus"
	"testing"
)

var dataVulgarWords = []struct {
	text string
	exp  float64
}{
	{
		"Canadian ", 1.0,
	}, {
		" (partition)|Bulkhead", 2.0,
	},
}

func TestVulgarWords(t *testing.T) {
	for _, td := range dataVulgarWords {
		got := tekstus.StringFrequenciesOf(td.text,
			tekstus.VulgarWords, false)

		assert(t, td.exp, got, true)
	}
}
