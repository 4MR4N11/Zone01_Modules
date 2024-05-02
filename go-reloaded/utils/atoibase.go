package reloaded

import "strings"

func valid(base string, s string) bool {
	tmp := []rune(base)
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(tmp); i++ {
		j := i + 1
		for j < len(tmp) {
			if tmp[i] == tmp[j] || tmp[i] == '-' || tmp[i] == '+' {
				return false
			}
			j++
		}
	}
	for _, a := range s {
		flag := false
		for _, b := range base {
			if a == b {
				flag = true
			}
		}

		if !flag {
			return false
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	res := 0
	if IsLower(s) {
		base = strings.ToLower(base)
	}
	if !valid(base, s) {
		return res
	}
	size := len(s) - 1
	for _, j := range s {
		index := strings.Index(base, string(j))
		res += index * RecursivePower(len(base), size)
		size--
	}
	return res
}
