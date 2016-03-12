// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

import (
	"fmt"
)

/*
Strings is for working with element of list with type is string.
Each element of slice is in the form of ["a", ..., "n"]
*/
type Strings []string

/*
Normalize return slice of string.
*/
func (ss *Strings) Normalize() []string {
	ls := make([]string, len(*ss))

	for x, v := range *ss {
		ls[x] = v
	}

	return ls
}

/*
IsEqual compare elements of two slice of string without regard to
their order

	{"a","b"} == {"b","a"} is true

Return true if each both slice have the same elements, false otherwise.
*/
func (ss *Strings) IsEqual(b Strings) bool {
	sslen := len(*ss)

	if sslen != len(b) {
		return false
	}

	check := make([]bool, sslen)

	for x, ls := range *ss {
		for _, rs := range b {
			if ls == rs {
				check[x] = true
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
StringsIsContain return true if elemen `el` is in slice of string `ss`, otherwise
return false.
*/
func StringsIsContain(ss Strings, el string) bool {
	for _, s := range ss {
		if s == el {
			return true
		}
	}
	return false
}

/*
SinglePartition create a table from a set of string, where each elemen in a set
become a single set.

Input: [a,b,c]
output:
    [
        [[a],[b],[c]]
    ]
*/
func (ss *Strings) SinglePartition() (table TableStrings) {
	list := make(ListStrings, len(*ss))

	for x, el := range *ss {
		list[x] = Strings{el}
	}

	table = append(table, list)
	return
}

/*
createIndent will create n space indentation and return it.
*/
func createIndent(n int) (s string) {
	for i := 0; i < n; i++ {
		s += " "
	}
	return
}

/*
Partitioning will group the set's element `orgseed` into non-empty
lists, in such a way that every element is included in one and only of the
lists.

Given a list of element in `orgseed`, and number of partition `k`, return
the set of all group of all elements without duplication.

For example, the set {a,b,c} if partitioned into 2 group will result in set

	{
		{{a,b},{c}},
		{{a,c},{b}},
		{{a},{b,c}},
	}

if partitioned into 3 group (k=3) will result in,

	{
		{{a},{b},{c}},
	}

Number of possible list can be computed using Stirling number of second kind.

For more information see,
- https://en.wikipedia.org/wiki/Partition_of_a_set
*/
func (ss *Strings) Partitioning(k int) (table TableStrings) {
	n := len(*ss)
	seed := make(Strings, n)
	copy(seed, *ss)

	if DEBUG >= 1 {
		fmt.Printf("[tekstus] %s Partitioning(%v,%v)\n", createIndent(n), n, k)
	}

	// if only one split return the set contain only seed as list.
	// input: {a,b,c},  output: {{a,b,c}}
	if k == 1 {
		list := make(ListStrings, 1)
		list[0] = seed

		table = append(table, list)
		return table
	}

	// if number of element in set equal with number split, return the set
	// that contain each element in list.
	// input: {a,b,c},  output:= {{a},{b},{c}}
	if n == k {
		return seed.SinglePartition()
	}

	// take the first element
	el := seed[0]

	// remove the first element from set
	seed = append(seed[:0], seed[1:]...)

	if DEBUG >= 1 {
		fmt.Printf("[tekstus] %s el: %s, seed:", createIndent(n), el, seed)
	}

	// generate child list
	genTable := seed.Partitioning(k)

	if DEBUG >= 1 {
		fmt.Printf("[tekstus] %s genTable join: %v", createIndent(n), genTable)
	}

	// join elemen with generated set
	table = genTable.JoinCombination(el)

	if DEBUG >= 1 {
		fmt.Printf("[tekstus] %s join %s      : %v\n", createIndent(n), el,
			table)
	}

	genTable = seed.Partitioning(k - 1)

	if DEBUG >= 1 {
		fmt.Printf("[tesktus] %s genTable append :", createIndent(n), genTable)
	}

	for _, lss := range genTable {
		list := make(ListStrings, len(lss))
		copy(list, lss)
		list = append(list, Strings{el})
		table = append(table, list)
	}

	if DEBUG >= 1 {
		fmt.Printf("[tesktus] %s append %v      : %v\n", createIndent(n), el,
			table)
	}

	return
}
