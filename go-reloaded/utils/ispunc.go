package reloaded

func IsPunc(s rune) bool {
	if s != ',' && s != '.' && s != ';' && s != '\'' && s != '?' && s != ':' && s != '!' {
		return false
	}
	return true
}
