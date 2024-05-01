package reloaded

import (
	"strconv"
	"strings"
)

func Converting(str string, op string) string {
	converted := ""
	if strings.TrimSpace(op) == "(hex)" {
		converted = strconv.Itoa(AtoiBase(strings.TrimSpace(str), "0123456789ABCDEF"))
	} else if strings.TrimSpace(op) == "(bin)" {
		converted = strconv.Itoa(AtoiBase(strings.TrimSpace(str), "01"))
	} else if strings.TrimSpace(op) == "(low)" {
		converted = strings.ToLower(str)
	} else if strings.TrimSpace(op) == "(up)" {
		converted = strings.ToUpper(str)
	} else if strings.TrimSpace(op) == "(cap)" {
		converted = Capitalize(str)
	} else if strings.Contains(op, "(low,") {
		converted = strings.ToLower(str)
	} else if strings.Contains(op, "(up,") {
		converted = strings.ToUpper(str)
	} else if strings.Contains(op, "(cap,") {
		converted = Capitalize(str)
	}
	if converted == "0" {
		return str
	}
	return converted
}
