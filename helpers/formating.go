package helpers

import "strings"

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
