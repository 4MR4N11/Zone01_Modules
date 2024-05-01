package reloaded

import (
	"fmt"
	"regexp"
	"strings"
)

func PunctHandler(data string) string {
	re := regexp.MustCompile(`[.,;:!?]\S`)
	regMatches := re.FindAllString(data, -1)
	fmt.Println(regMatches)
	for _, m := range regMatches {
		index := strings.Index(data, m)
		if index != -1 {
			if index+1 < len(data) {
				data = data[:index+1] + " " + data[index+1:]
			}
		}
	}
	fmt.Println(data)
	re = regexp.MustCompile(`\s+[.,;:!?]`)
	regMatches = re.FindAllString(data, -1)
	for _, m := range regMatches {
		index := strings.Index(data, m)
		if index != -1 {
			data = data[:index] + data[len(m)+index-1:]
		}
	}
	return data
}
