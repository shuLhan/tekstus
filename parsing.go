// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tekstus

/*
ParsingSkipUntil skip all bytes until matched token is found.

If `checkEsc` is true, token that is prefixed with escaped character
'\' will be considered as non-match token.

Return index of line with matched token or false if line end before
finding the token.
*/
func ParsingSkipUntil(token, line []byte, startAt int, checkEsc bool) (
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
		found = TokenMatchForward(token, line, p)

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
ParsingUntil we found token.

If `checkEsc` is true, token that is prefixed with escaped character
'\' will be considered as non-match token.

Return all bytes before token and positition of byte _after_ token,
or false if token is not found.
*/
func ParsingUntil(token, line []byte, startAt int, checkEsc bool) (
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
		found = TokenMatchForward(token, line, p)

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
