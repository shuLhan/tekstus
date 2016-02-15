// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

/*
IsIn return true if character `c` is in `s`.
*/
func IsIn(s []rune, c rune) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}
	return false
}
