package reloaded

import (
	"regexp"
	"strings"
)

func PunctHandler(data string) string {
	reg := regexp.MustCompile(`[.,;:!?]\S`)
	regMatches := reg.FindAllString(data, -1)
	for _, m := range regMatches {
		index := strings.Index(data, m)
		if index != -1 {
			if index+1 < len(data) {
				data = data[:index+1] + " " + data[index+1:]
			}
		}
	}
	reg = regexp.MustCompile(`\s+[.,;:!?]`)
	regMatches = reg.FindAllString(data, -1)
	for _, m := range regMatches {
		index := strings.Index(data, m)
		if index != -1 {
			data = data[:index] + data[len(m)+index-1:]
		}
	}
	return data
}
