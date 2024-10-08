package reloaded

import (
	"strings"
	"unicode"
)

// this functions uppers the word before the flag and returns true if it works
func Up(sl []string, index int) ([]string, bool) {
	works := false
	for i := index - 1; i >= 0; i-- {
		if isAlpha(sl[i]) {
			works = true
			slice := []rune(sl[i])
			for j := 0; j < len(slice); j++ {
				if unicode.IsLetter(slice[j]) {
					slice[j] = unicode.ToUpper(slice[j])
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

// Capitalize first n alphabetic words before index
func UpNumbers(sl []string, index, n int) []string {
	for i := index - 1; i >= 0 && n > 0; i-- {
		if isAlpha(sl[i]) {
			sl[i] = strings.ToUpper(sl[i])
			n--
		}
	}
	return sl
}
