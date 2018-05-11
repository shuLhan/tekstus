// Copyright 2016-2018 Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

//
// TableStrings is for working with set of list of set of string.
// Each elemen in set is in the form of
//
//	[
//		[["a"],["b","c"],...],
//		[["x"],["y",z"],...]
//	]
//
type TableStrings []ListStrings

//
// IsEqual compare two table of string without regard to their order.
//
//	{
//		{{"a"},{"b"}},
//		{{"c"}}
//	}
//
// is equal to
//
//	{
//		{{"c"}},
//		{{"b"},{"a"}}
//	}
//
// Return true if both set is contain the same list, false otherwise.
//
func (tss *TableStrings) IsEqual(b TableStrings) bool {
	tsslen := len(*tss)

	if tsslen != len(b) {
		return false
	}

	check := make([]bool, tsslen)

	for x, llss := range *tss {
		for _, rlss := range b {
			if llss.IsEqual(rlss) {
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

//
// JoinCombination will append string `s` to each set in list in different index.
//
// For example, given string `s` and input table `[[["a"]["b"]["c"]]]`, the
// output table will be,
//
//	[
//		[["a","s"]["b"]    ["c"]],
//		[["a"]    ["b","s"]["c"]],
//		[["a"]    ["b"]    ["c","s"]]
//	]
//
func (tss *TableStrings) JoinCombination(s string) (tssout TableStrings) {
	for _, lss := range *tss {
		for y := range lss {
			newList := make(ListStrings, len(lss))
			copy(newList, lss)
			newList[y] = append(newList[y], s)
			tssout = append(tssout, newList)
		}
	}
	return
}
