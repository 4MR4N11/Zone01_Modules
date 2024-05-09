package reloaded

import "strings"

func isVowel(data rune) bool {
	vowels := "aeiouhAEIOUH"
	for _, v := range vowels {
		if data == v {
			return true
		}
	}
	return false
}

func VowelHandler(data string) string {
	tmp := strings.Split(data, " ")
	for i := 0; i < len(tmp); i++ {
		if tmp[i] == "a" || tmp[i] == "A" {
			j := i + 1
			for ; j < len(tmp) && len(tmp[j]) == 0; j++ {
			}
			if j < len(tmp) && len(tmp[j]) > 0 && isVowel(rune(tmp[j][0])) {
				tmp[i] += "n"
			}
		}
	}
	data = strings.Join(tmp, " ")
	return data
}
