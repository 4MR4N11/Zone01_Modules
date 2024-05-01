package reloaded

import (
	"fmt"
	"regexp"
	"strings"
)

func Quotehandler(data string) string {
	re := regexp.MustCompile(`\s+'|'\s+`)
	regMatches := re.FindAllString(data, -1)
	if len(regMatches) > 0 && len(regMatches)%2 != 0 {
		regMatches = regMatches[:len(regMatches)-1]
	}
	for _, m := range regMatches {
		quote := 1
		index := strings.Index(data, m)
		// handle every 2 quotes at a time
	}
	return data
}
