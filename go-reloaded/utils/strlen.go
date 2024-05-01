package reloaded

func StrLen(s string) int {
	str := []rune(s)
	len := 0
	for j := range str {
		j++
		len++
	}
	return len
}
