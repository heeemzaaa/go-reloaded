package reloaded

import "strconv"

// this function checks if the string passed is a hexadecimal number
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

// this function convert the hexadecimal number to a decimal
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
