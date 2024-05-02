package reloaded

func runeIsAlphaNum(char rune) bool {
	if char < 'a' || char > 'z' {
		if char < 'A' || char > 'Z' {
			if char < '0' || char > '9' {
				return false
			}
		}
	}
	return true
}

func runeIsAlpha(char rune) bool {
	if (char <= 'z' && char >= 'a') || (char <= 'Z' && char >= 'A') {
		return true
	}
	return false
}

func runeIsLower(char rune) bool {
	if char < 'a' || char > 'z' {
		return false
	}
	return true
}

func runeToUpper(char rune) rune {
	if runeIsLower(char) {
		char -= 32
	}
	return char
}

func Capitalize(s string) string {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if runeIsAlpha(str[i]) {
			if i == 0 {
				str[i] = runeToUpper(str[i])
			} else if i-2 >= 0 && !runeIsAlphaNum(str[i-1]) {
				if (str[i-1] == '\'' && str[i-2] == ' ') || str[i-1] == ' ' {
					str[i] = runeToUpper(str[i])
				}
			} else if !runeIsLower(str[i]) {
				str[i] += 32
			}
		}
	}
	return string(str)
}
