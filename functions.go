package reloaded

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

// function to capitalize the first letter in the word
func HandleTheFlags(line []string) []string {
	for i := 0; i < len(line); i++ {
		if line[i] == "(cap)" {
			if i-1 >= 0 {
				_, b := Cap(line, i)
				if !b {
					log.Println("The flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					line = append(line[:i], line[i+1:]...)
					i = 0
				}
			}
		} else if line[i] == "(up)" {
			if i-1 >= 0 {
				_, b := Up(line, i)
				if !b {
					log.Println("The flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					line = append(line[:i], line[i+1:]...)
					i = 0
				}
			}
		} else if line[i] == "(low)" {
			if i-1 >= 0 {
				_, b := Low(line, i)
				if !b {
					log.Println("The flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					line = append(line[:i], line[i+1:]...)
					i = 0
				}
			}
		} else if line[i] == "(cap," {
			if i+1 < len(line) && line[i+1] != "" {
				n, _ := strconv.Atoi(line[i+1][:len(line[i+1])-1])
				if n == 0 || n > i {
					log.Println("The flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					CapNumbers(line, i, n)
					line = append(line[:i], line[i+2:]...)
					i = 0
				}
			}
		} else if line[i] == "(up," {
			if i+1 < len(line) && line[i+1] != "" {
				n, _ := strconv.Atoi(line[i+1][:len(line[i+1])-1])
				if n == 0 || n > i {
					log.Println("The flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					UpNumbers(line, i, n)
					line = append(line[:i], line[i+2:]...)
					i = 0
				}
			}
		} else if line[i] == "(low," {
			if i+1 < len(line) && line[i+1] != "" {
				n, _ := strconv.Atoi(line[i+1][:len(line[i+1])-1])
				if n == 0 || n > i {
					log.Println("The (low,) flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					LowNumbers(line, i, n)
					line = append(line[:i], line[i+2:]...)
					i = 0
				}
			}
		} else if line[i] == "(hex)" {
			if i-1 >= 0 {
				_, b := HexToDecimal(line, i)
				if !b {
					log.Println("The flag (hex) doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					line = append(line[:i], line[i+1:]...)
					i = 0
				}
			}
		} else if line[i] == "(bin)" {
			if i-1 >= 0 {
				_, b := BinToDecimal(line, i)
				if !b {
					log.Println("The (bin) flag doesn't work correctly so it would be handled as a string !")
					i++
				} else {
					line = append(line[:i], line[i+1:]...)
					i = 0
				}
			}
		}
	}
	return line
}

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

// this function is to capitalize a number of strings
func CapNumbers(sl []string, index, n int) []string {
	for i := index - 1; i >= 0 && n > 0; i-- {
		if isAlpha(sl[i]) {
			sl[i] = strings.Title(sl[i])
			n--
		}
	}
	return sl
}

func LowNumbers(sl []string, index, n int) []string {
	for i := index - 1; i >= 0 && n > 0; i-- {
		if isAlpha(sl[i]) {
			sl[i] = strings.ToLower(sl[i])
			n--
		}
	}
	return sl
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

func IsPunctuation(r rune) bool {
	return r == ',' || r == '.' || r == ';' || r == ':' || r == '?' || r == '!'
}

func SplitWhiteSpaces(s string) []string {
	slice := []rune(s)
	holder := ""
	result := []string{}

	for i := 0; i < len(slice); i++ {
		if slice[i] != ' ' {
			holder += string(slice[i])
		} else if slice[i] == ' ' {
			if holder != "" {
				result = append(result, holder)
				holder = ""
			}
		}
	}
	if holder != "" {
		result = append(result, holder)
		holder = ""
	}
	return result
}

func HandleQuotes(text string) string {
	statut := false
	var result []rune
	slice := []rune(text)
	for i := 0; i < len(slice); i++ {
		if slice[i] == '\'' && (i+1 < len(slice) && i-1 >= 0) && (unicode.IsLetter(slice[i+1]) && unicode.IsLetter(slice[i-1])) {
			result = append(result, rune(slice[i]))
			i++
		}
		if slice[i] == '\'' {
			if !statut {
				statut = true
				if i-1 >= 0 && slice[i-1] != ' ' && len(result) > 0 {
					result = append(result, ' ')
				}
				result = append(result, rune(slice[i]))
				if i+1 < len(slice) && slice[i+1] == ' ' {
					i++
				}
			} else {
				statut = false
				if len(result) > 0 && result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result = append(result, rune(slice[i]))
				if i+1 < len(slice) && slice[i+1] != ' ' {
					result = append(result, ' ')
				}

			}
		} else {
			result = append(result, rune(slice[i]))
		}
	}
	return string(result)
}

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

func IsHex(s string) bool {
	result := true
	sl := []rune(s)
	for i := 0; i < len(sl); i++ {
		if (sl[i] >= '0' && sl[i] <= '9') || (sl[i] >= 'A' && sl[i] <= 'F') || (sl[i] >= 'a' && sl[i] <= 'f') {
			result = true
		} else {
			result = false
			break
		}
	}

	return result
}

func IsBin(s string) bool {
	result := true
	sl := []rune(s)
	for i := 0; i < len(sl); i++ {
		if sl[i] == '0' || sl[i] == '1' {
			result = true
		} else {
			result = false
			break
		}
	}

	return result
}

func HexToDecimal(sl []string, index int) ([]string, bool) {
	works := false
	for i := index - 1; i >= 0; i-- {
		if IsHex(sl[i]) {
			works = true
			decimal, err := strconv.ParseInt(sl[i], 16, 64)
			if err != nil {
				works = false
				return sl, works
			}
			sl[i] = strconv.Itoa(int(decimal))
			break
		} else if i == 0 && works == false {
			return sl, works
		}
	}

	return sl, works
}

func BinToDecimal(sl []string, index int) ([]string, bool) {
	works := false
	for i := index - 1; i >= 0; i-- {
		if IsBin(sl[i]) {
			works = true
			decimal, err := strconv.ParseInt(sl[i], 2, 64)
			if err != nil {
				works = false
				return sl, works
			}
			sl[i] = strconv.Itoa(int(decimal))
			break
		} else if i == 0 && works == false {
			return sl, works
		}
	}

	return sl, works
}
