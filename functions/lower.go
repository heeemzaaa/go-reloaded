package reloaded

import (
	"strings"
	"unicode"
)

// this function lowers the word before and return also a true if it works
func Low(sl []string, index int) ([]string, bool) {
	works := false
	for i := index - 1; i >= 0; i-- {
		if isAlpha(sl[i]) {
			works = true
			slice := []rune(sl[i])
			for j := 0; j < len(slice); j++ {
				if unicode.IsLetter(slice[j]) {
					slice[j] = unicode.ToLower(slice[j])
				}
			}
			sl[i] = string(slice)
			break
		} else if i == 0 && works == false {
			return sl, works
		}
	}
	return sl, works
}

// this function lowers a number of words
func LowNumbers(sl []string, index, n int) []string {
	for i := index - 1; i >= 0 && n > 0; i-- {
		if isAlpha(sl[i]) {
			sl[i] = strings.ToLower(sl[i])
			n--
		}
	}
	return sl
}
