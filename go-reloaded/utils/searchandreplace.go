package reloaded

import (
	"regexp"
	"strings"
)

func validOp(op string) bool {
	if strings.TrimSpace(op) == "(hex)" || strings.TrimSpace(op) == "(bin)" || strings.TrimSpace(op) == "(low)" || strings.TrimSpace(op) == "(up)" || strings.TrimSpace(op) == "(cap)" || strings.Contains(op, "(low,") || strings.Contains(op, "(up,") || strings.Contains(op, "(cap,") {
		return true
	}
	return false
}

func SearchAndReplaceOp(str string) string {
	re := regexp.MustCompile(`\((\w+)(?:,\s(\d*[1-9]\d*))?\)`)
	regTmp := re.FindAllString(str, -1)
	regMatches := []string{}
	for _, m := range regTmp {
		if validOp(m) {
			regMatches = append(regMatches, m)
		}
	}
	for _, m := range regMatches {
		if strings.HasSuffix(str, " "+m) || strings.HasPrefix(str, m+" ") || strings.Contains(str, " "+m+" ") {
			if strings.HasSuffix(str, " "+m) {
				m = " " + m
				match := Getmatch(str, m)
				str = strings.ReplaceAll(str, match+m, Converting(match, strings.TrimSpace(m)))
			} else if strings.HasPrefix(str, m+" ") {
				m = m + " "
				str = strings.Trim(str, m)
			} else {
				m = " " + m
				match := Getmatch(str, m)
				str = strings.ReplaceAll(str, match+m, Converting(match, strings.TrimSpace(m)))
			}
		} else if str == m {
			str = ""
		}
	}
	return str
}
