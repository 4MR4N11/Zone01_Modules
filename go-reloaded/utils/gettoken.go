package reloaded

import (
	"strconv"
	"strings"
)

func notSpace(str string) bool {
	for _, char := range str {
		if char != ' ' {
			return true
		}
	}
	return false
}

func Getmatch(str string, op string) string {
	match := ""
	index := strings.Index(str, op)
	if index != -1 {
		if strings.Contains(op, ", ") {
			tmp := strings.Split(strings.TrimSpace(op), ", ")
			if tmp[0] == "(up" || tmp[0] == "(low" || tmp[0] == "(cap" {
				num, err := strconv.Atoi(strings.Trim(tmp[1], ")"))
				if err == nil {
					for i := index - 1; i >= 0; i-- {
						if str[i] == ' ' && len(strings.Fields(match)) == num {
							break
						}
						match = string(str[i]) + match
					}
				}
			}
		} else {
			for i := index - 1; i >= 0; i-- {
				if str[i] == ' ' && len(match) > 0 && notSpace(match) {
					break
				}
				match = string(str[i]) + match
			}
		}
	}
	return match
}
