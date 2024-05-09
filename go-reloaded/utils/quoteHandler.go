package reloaded

import (
	"strings"
)

func Quotehandler(data string) string {
	tmp := ""
	for i := 0; i < len(data); i++ {
		if i+1 < len(data) && i-1 >= 0 && data[i] == '\'' && (data[i-1] == ' ' || data[i+1] == ' ' || i == 0 || i == len(data)-1) {
			j := i + 1
			for ; j < len(data); j++ {
				if data[j] == '\'' {
					if j+1 < len(data) && data[j+1] == ' ' {
						break
					} else if j-1 >= 0 && data[j-1] == ' ' {
						break
					} else if j+1 < len(data) && !IsAlpha(string(data[j+1])) {
						break
					}
				} else if j+1 == len(data) && data[j] != '\'' {
					tmp = ""
					break
				}
				tmp += string(data[j])
			}
			if tmp != "" {
				data = data[:i+1] + strings.TrimSpace(tmp) + data[j:]
				i += len(strings.TrimSpace(tmp)) + 1
				tmp = ""
			}
		}
	}
	return data
}
