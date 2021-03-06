// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

import (
	"fmt"
)

//
// InterfacesToStrings will convert slice of interface to slice of string.
//
func InterfacesToStrings(is []interface{}) (vs []string) {
	for _, i := range is {
		v := fmt.Sprintf("%v", i)
		vs = append(vs, v)
	}
	return
}
