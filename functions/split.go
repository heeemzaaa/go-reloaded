package reloaded

// this function splits our string with spaces to a slice of string
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
