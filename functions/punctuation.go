package reloaded

import "strings"

// this function handles all the punctuation cases
func HandlePunctuations(s string) string {
	slice := SplitWhiteSpaces(s)

	for i := 0; i < len(slice); i++ {
		if i+1 < len(slice) && IsPunctuation(rune(slice[i+1][0])) {
			slice[i] = slice[i] + string(slice[i+1][0])
			slice[i+1] = slice[i+1][1:]
			s = strings.Join(slice, " ")
			slice = SplitWhiteSpaces(s)
			i = -1
		}
	}
	return strings.Join(slice, " ")
}

// this function defines the punctuation that we need to handle
func IsPunctuation(r rune) bool {
	return r == ',' || r == '.' || r == ';' || r == ':' || r == '?' || r == '!'
}
