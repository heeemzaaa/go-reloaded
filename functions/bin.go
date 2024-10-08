package reloaded

import "strconv"

// this function checks if the string passed is a binary number
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


// this function converts a binary number to a decimal one
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
