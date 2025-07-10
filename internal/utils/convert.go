package utils

import "strings"

func NomalizeText(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}
