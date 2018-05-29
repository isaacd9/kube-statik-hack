package snake

import (
	"bytes"
)

func Snake(s string) string {
	ret := &bytes.Buffer{}

	for i := 0; i < len(s); {
		c := s[i]

		if isUpper(c) {
			if i != 0 {
				ret.WriteRune('_')
			}
			ret.WriteByte(toLower(c))
		} else if isDigit(c) {
			if i != 0 {
				ret.WriteRune('_')
			}

			// Append rest of digits in sequence
			for {
				c = s[i]
				if !isDigit(c) || i > len(s) {
					i--
					break
				}
				ret.WriteByte(c)
				i++
			}
		} else {
			ret.WriteByte(c)
		}
		i++
	}

	return ret.String()
}

func isWord(c byte) bool {
	return isLetter(c) || isDigit(c)
}

func isLetter(c byte) bool {
	return isLower(c) || isUpper(c)
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func toLower(c byte) byte {
	if isUpper(c) {
		return c + ('a' - 'A')
	}
	return c
}

func toUpper(c byte) byte {
	if isLower(c) {
		return c - ('a' - 'A')
	}
	return c
}
