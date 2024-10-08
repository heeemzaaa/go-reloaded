package reloaded

// this function handle all the vowel situations
func VowelSituation(sl []string) []string {
	for i := 0; i < len(sl); i++ {
		if sl[i] == "a" || sl[i] == "A" {
			if i+1 < len(sl) {
				nextWord := sl[i+1]

				if len(nextWord) > 0 && isVowel(rune(nextWord[0])) {
					if sl[i] == "a" {
						sl[i] = "an"
					} else {
						sl[i] = "An"
					}
				}
			}
		}
	}
	return sl
}

// Helper function to check if a character is a vowel
func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' || r == 'H'
}
