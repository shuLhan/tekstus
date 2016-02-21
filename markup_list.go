// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

// WikiMarkup define the markup for Wikimedia software.
type WikiMarkup struct {
	begin string
	end   string
}

// WikiMarkups contain list of common markup in Wikimedia software.
var WikiMarkups = []WikiMarkup{
	{
		"[[Category:",
		"]]",
	}, {
		"[[:Category:",
		"]]",
	}, {
		"[[File:",
		"]]",
	}, {
		"[[Help:",
		"]]",
	}, {
		"[[Image:",
		"]]",
	}, {
		"[[Special:",
		"]]",
	}, {
		"[[Wikipedia:",
		"]]",
	}, {
		"{{DEFAULTSORT:",
		"}}",
	}, {
		"{{Template:",
		"}}",
	}, {
		"<ref",
		"/>",
	},
}
