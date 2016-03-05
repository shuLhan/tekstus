// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package tekstus is a library for working with text.
*/
package tekstus

import (
	"os"
	"strconv"
)

const (
	// DefEscape character(s).
	DefEscape = '\\'
)

var (
	// DEBUG debug level, set using environment TEKSTUS_DEBUG
	DEBUG = 0
)

func init() {
	v := os.Getenv("TEKSTUS_DEBUG")
	if v == "" {
		DEBUG = 0
	} else {
		DEBUG, _ = strconv.Atoi(v)
	}
}
