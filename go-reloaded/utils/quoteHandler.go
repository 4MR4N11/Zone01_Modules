package reloaded

import (
	"strings"
)

func Quotehandler(data string) string {
	tmp := ""
	for i := 0; i < len(data); i++ {
		if data[i] == '\'' {
			j := i + 1
			for ; j < len(data); j++ {
				if data[j] == '\'' {
					if j+1 < len(data) && data[j+1] == ' ' {
						break
					} else if j-1 >= 0 && data[j-1] == ' ' {
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
