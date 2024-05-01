package reloaded

func StrRev(s string) string {
	tmp := []byte(s)
	str := []byte(s)
	len := StrLen(s) - 1
	for i := 0; len >= 0; i++ {
		str[i] = tmp[len]
		len--
	}
	return string(str)
}
