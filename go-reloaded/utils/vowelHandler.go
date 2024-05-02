package reloaded

import (
	"regexp"
	"strings"
)

func VowelHandler(data string) string {
	re := regexp.MustCompile(`\b[aA]\s+[aeiouhAEIOUH]`)
	regMatches := re.FindAllString(data, -1)
	for _, m := range regMatches {
		if m[0] == 'a' {
			data = strings.ReplaceAll(data, m, "an "+m[2:])
		} else {
			data = strings.ReplaceAll(data, m, "An "+m[2:])
		}
	}
	return data
}
