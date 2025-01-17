package utils

import "strings"

func CleanInput(s string) []string {
	clean := strings.Fields(strings.ToLower(s))
	return clean
}
