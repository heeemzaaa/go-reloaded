package reloaded

import "unicode"

// this function checks if the string passed is alphabetical or not
func isAlpha(s string) bool {
	// Check if the string contains at least one letter
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
