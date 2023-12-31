package helpers

import (
	"regexp"
	"strings"
)

func FormatLocationText(original string) string {
	// auckland-new_zealand --> Auckland - New Zealand
	result := ""
	for i, location := range strings.Split(original, "-") {
		for j, word := range strings.Split(location, "_") {
			if j < len(strings.Split(location, "_"))-1 {
				result += strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:])) + " "
			} else {
				result += strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
			}
		}
		if i < len(strings.Split(original, "-"))-1 {
			result += " - "
		}
	}
	return result
}

func FormatLocationMaps(original string) string {

	locationWithoutDash := strings.ReplaceAll(original, "-", "")

	// Replace spaces with '+' in the formatted string
	re := regexp.MustCompile(`\s+`)
	formattedLocation := re.ReplaceAllString(locationWithoutDash, "+")

	return formattedLocation
}
