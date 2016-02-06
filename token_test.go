// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus_test

import (
	"github.com/shuLhan/tekstus"
	"reflect"
	"runtime/debug"
	"testing"
)

func assert(t *testing.T, exp, got interface{}, equal bool) {
	if reflect.DeepEqual(exp, got) != equal {
		debug.PrintStack()
		t.Fatalf("Expecting '%v' got '%v'\n", exp, got)
	}
}

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
	line := []byte("// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.")

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
	line := []byte("// Copyright 2016 Mhd Sulhan \"<ms@kilabit.info>\". All rights reserved.")

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
