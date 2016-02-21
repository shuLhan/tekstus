// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	_ "fmt"
	"github.com/shuLhan/tekstus"
	"testing"
)

func testFindToken(t *testing.T, token, line []byte, startat int, exp []int) {
	got := []int{}
	tokenlen := len(token)

	for {
		foundat := tekstus.FindToken(token, line, startat)

		if foundat < 0 {
			break
		}

		got = append(got, foundat)
		startat = foundat + tokenlen
	}

	assert(t, exp, got, true)
}

func TestFindToken(t *testing.T) {
	line := []byte(dataLines[0])

	token := []byte("//")
	exp := []int{0}

	testFindToken(t, token, line, 0, exp)

	token = []byte(".")
	exp = []int{40, 46, 67}

	testFindToken(t, token, line, 0, exp)

	token = []byte("d.")
	exp = []int{66}

	testFindToken(t, token, line, 0, exp)
}

func testEncasulateToken(t *testing.T, token, line, leftcap, rightcap,
	exp []byte) {

	newline, changed := tekstus.EncapsulateToken(token, line, leftcap, rightcap)

	assert(t, string(exp), string(newline), changed)
}

func TestEncapsulateToken(t *testing.T) {
	line := []byte(dataLines[1])

	token := []byte("/")
	leftcap := []byte("\\")
	rightcap := []byte{}
	exp := []byte("\\/\\/ Copyright 2016 Mhd Sulhan \"<ms@kilabit.info>\". All rights reserved.")

	testEncasulateToken(t, token, line, leftcap, rightcap, exp)

	token = []byte("<")
	leftcap = []byte("<")
	rightcap = []byte(" ")
	exp = []byte("// Copyright 2016 Mhd Sulhan \"<< ms@kilabit.info>\". All rights reserved.")

	testEncasulateToken(t, token, line, leftcap, rightcap, exp)

	token = []byte("\"")
	leftcap = []byte("\\")
	rightcap = []byte(" ")
	exp = []byte("// Copyright 2016 Mhd Sulhan \\\" <ms@kilabit.info>\\\" . All rights reserved.")

	testEncasulateToken(t, token, line, leftcap, rightcap, exp)
}

func TestEncapsulateTrim(t *testing.T) {
	for _, td := range dataEncapsulateTrim {
		got, _ := tekstus.EncapsulateTrim([]byte(td.text),
			[]byte(td.leftcap), []byte(td.rightcap))

		assert(t, string(td.exp), string(got), true)
	}
}
