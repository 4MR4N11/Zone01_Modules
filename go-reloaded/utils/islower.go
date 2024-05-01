package reloaded

func IsLower(s string) bool {
	for _, letter := range s {
		if letter < 'a' || letter > 'z' {
			if letter < '0' || letter > '9' {
				if letter != ' ' {
					return false
				}
			}
		}
	}
	return true
}
