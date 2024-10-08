package reloaded

import (
	"strings"
	"unicode"
)

// this function is for capitalizing the first charcter in the first string before (cap)
func Cap(sl []string, index int) ([]string, bool) {
	count := 0
	works := false
	for i := index - 1; i >= 0; i-- {
		if isAlpha(sl[i]) {
			works = true
			slice := []rune(sl[i])
			for j := 0; j < len(slice); j++ {
				if unicode.IsLetter(slice[j]) && count > 0 {
					slice[j] = unicode.ToLower(slice[j])
				} else if unicode.IsLetter(slice[j]) {
					slice[j] = unicode.ToUpper(slice[j])
					count++
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

// this function capitalizes a number of strings
func CapNumbers(sl []string, index, n int) []string {
	for i := index - 1; i >= 0 && n > 0; i-- {
		if isAlpha(sl[i]) {
			sl[i] = strings.Title(sl[i])
			n--
		}
	}
	return sl
}
