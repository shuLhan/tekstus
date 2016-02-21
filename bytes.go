// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

/*
BytesFind return the first index of matched token in line.
If not found it will return -1.
*/
func BytesFind(line, token []byte, startat int) (at int) {
	y := 0
	tokenlen := len(token)
	linelen := len(line)

	at = -1
	for x := startat; x < linelen; x++ {
		if line[x] == token[y] {
			if y == 0 {
				at = x
			}
			y++
			if y == tokenlen {
				// we found it!
				return
			}
		} else {
			if at != -1 {
				// reset back
				y = 0
				at = -1
			}
		}
	}
	// x run out before y
	if y < tokenlen {
		at = -1
	}
	return
}

/*
BytesMatchForward return true if `line` at index `p` match with `token`,
otherwise return false.
*/
func BytesMatchForward(line, token []byte, p int) bool {
	linelen := len(line)
	tokenlen := len(token)

	if p+tokenlen > linelen {
		return false
	}

	for _, v := range token {
		if v != line[p] {
			return false
		}
		p++
	}
	return true
}

/*
BytesEncapsulate will find `token` in `line` and capsulating it with bytes
from `leftcap` and `rightcap`.
If no token is found, it will return the same line with false status.
*/
func BytesEncapsulate(token, line, leftcap, rightcap []byte) (
	newline []byte,
	changed bool,
) {
	tokenlen := len(token)

	startat := 0
	for {
		foundat := BytesFind(line, token, startat)

		if foundat < 0 {
			newline = append(newline, line[startat:]...)
			break
		}

		newline = append(newline, line[startat:foundat]...)
		newline = append(newline, leftcap...)
		newline = append(newline, token...)
		newline = append(newline, rightcap...)

		startat = foundat + tokenlen
	}

	if startat > 0 {
		changed = true
	}

	return
}

/*
BytesRemoveUntil given a line, remove all bytes inside it, starting from
`leftcap` until the `rightcap` and return cutted line and changed to true.

If no `leftcap` or `rightcap` is found, the line will unchanged, and changed
will be false.

Example,

	line    : "[[ ABC ]] DEF"
	leftcap : "[["
	rightcap: "]]"
	return  : "  DEF"
*/
func BytesRemoveUntil(line, leftcap, rightcap []byte) (
	newline []byte,
	changed bool,
) {
	lidx := BytesFind(line, leftcap, 0)
	ridx := BytesFind(line, rightcap, lidx+1)

	if lidx < 0 || ridx < 0 || lidx >= ridx {
		return line, false
	}

	newline = line[:lidx]
	newline = append(newline, line[ridx+len(rightcap):]...)
	changed = true

	// Repeat
	newline, _ = BytesRemoveUntil(newline, leftcap, rightcap)

	return
}

/*
BytesSkipUntil skip all bytes until matched token is found.

If `checkEsc` is true, token that is prefixed with escaped character
'\' will be considered as non-match token.

Return index of line with matched token or false if line end before
finding the token.
*/
func BytesSkipUntil(line, token []byte, startAt int, checkEsc bool) (
	p int,
	found bool,
) {
	linelen := len(line)
	escaped := false

	for p = startAt; p < linelen; p++ {
		// Check if the escape character is used to escaped the
		// token ...
		if checkEsc && line[p] == DefEscape {
			escaped = true
			continue
		}
		if line[p] != token[0] {
			goto pass
		}

		// We found the first token character.
		// Lets check if its match with all content of token.
		found = BytesMatchForward(line, token, p)

		// False alarm ...
		if !found {
			goto pass
		}

		// Its matched, but if its prefixed with escaped char, then
		// we assumed it as non breaking token.
		if checkEsc && escaped {
			escaped = false
			continue
		}

		// We found the token at `p`
		p = p + len(token)
		found = true
		break

	pass:
		if escaped {
			// ... turn out its not escaping token.
			escaped = false
		}
	}

	return p, found
}

/*
BytesCutUntil we found token.

If `checkEsc` is true, token that is prefixed with escaped character
'\' will be considered as non-match token.

Return all bytes before token and positition of byte _after_ token,
or false if token is not found.
*/
func BytesCutUntil(line, token []byte, startAt int, checkEsc bool) (
	v []byte,
	p int,
	found bool,
) {
	linelen := len(line)
	tokenlen := len(token)
	escaped := false

	for p = startAt; p < linelen; p++ {
		// Check if the escape character is used to escaped the
		// token ...
		if checkEsc && line[p] == DefEscape {
			if escaped {
				// escaped already, its mean double '\\'
				v = append(v, '\\')
				escaped = false
			} else {
				escaped = true
			}
			continue
		}
		if line[p] != token[0] {
			goto pass
		}

		// We found the first token character.
		// Lets check if its match with all content of token.
		found = BytesMatchForward(line, token, p)

		// False alarm ...
		if !found {
			goto pass
		}

		// Found it, but if its prefixed with escaped char, then
		// we assumed it as non breaking token.
		if checkEsc && escaped {
			v = append(v, token...)
			p = p + tokenlen - 1
			escaped = false
			continue
		}

		// We found the token match in `line` at `p`
		return v, p + tokenlen, true

	pass:
		if escaped {
			// ... turn out its not escaping token.
			v = append(v, DefEscape)
			escaped = false
		}
		v = append(v, line[p])
	}

	// We did not found it...
	return v, p, false
}
