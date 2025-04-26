package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func removeSpaces(str string) string {
	splited := strings.Split(str, " ")
	var formated string

	for i := range splited {
		if splited[i] != "" {
			formated += strings.TrimSpace(splited[i])
			if i+1 < len(splited) {
				formated += " "
			}
		}
	}

	return formated
}

func CleanString(str string) string {
	str = removeSpaces(strings.ToUpper(strings.TrimSpace(str)))
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, str)
	return result
}
