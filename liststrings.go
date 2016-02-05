// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

import (
	"strings"
)

/*
ListStrings is for working with list of set of string.
Each elemen of slice is in the form of [["a"],["b","c"],...]
*/
type ListStrings []Strings

/*
IsEqual compare two list of slice of string without regard to
their order.

	{{"a"},{"b"}} == {{"b"},{"a"}} is true.

Return true if both contain the same list, false otherwise.
*/
func (lss *ListStrings) IsEqual(b ListStrings) bool {
	lsslen := len(*lss)

	if lsslen != len(b) {
		return false
	}

	check := make([]bool, lsslen)

	for x, lss := range *lss {
		for _, rstrings := range b {
			if lss.IsEqual(rstrings) {
				check[x] = true
				break
			}
		}
	}

	for _, v := range check {
		if !v {
			return false
		}
	}
	return true
}

/*
Join list of slice of string using `lsep` as separator between list items and
`ssep` for element in each slice.
*/
func (lss *ListStrings) Join(lsep string, ssep string) (s string) {
	lsslen := len(*lss) - 1

	for x, ls := range *lss {
		s += strings.Join(ls, ssep)

		if x < lsslen {
			s += lsep
		}
	}
	return
}
