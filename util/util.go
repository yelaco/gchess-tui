package util

import "strings"

func GenerateBlankLine(count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat("\n", count)
}
