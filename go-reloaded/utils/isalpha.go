package reloaded

func IsAlpha(s string) bool {
	if StrLen(s) == 0 {
		return true
	}
	for _, letter := range s {
		if letter < 'a' || letter > 'z' {
			if letter < 'A' || letter > 'Z' {
				if letter < '0' || letter > '9' {
					return false
				}
			}
		}
	}
	return true
}
