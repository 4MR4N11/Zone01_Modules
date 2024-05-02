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
	tmp := []rune(s)
	for i := 0; i < len(tmp); i++ {
		if runeIsAlpha(tmp[i]) {
			if i == 0 {
				tmp[i] = runeToUpper(tmp[i])
			} else if i-2 >= 0 && !runeIsAlphaNum(tmp[i-1]) {
				if (tmp[i-1] == '\'' && tmp[i-2] == ' ') || tmp[i-1] == ' ' {
					tmp[i] = runeToUpper(tmp[i])
				}
			} else if !runeIsLower(tmp[i]) {
				tmp[i] += 32
			}
		}
	}
	return string(tmp)
}
