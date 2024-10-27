package util

import "strings"

func GenerateBlankLine(count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat("\n", count)
}

func BoardToFen(board [][]string) string {
	var fen strings.Builder
	for _, row := range board {
		empty := 0
		for _, cell := range row {
			if cell == "" {
				empty++
			} else {
				if empty > 0 {
					fen.WriteString(string(rune(empty + 48)))
					empty = 0
				}
				fen.WriteString(cell)
			}
		}
		if empty > 0 {
			fen.WriteString(string(rune(empty + 48)))
		}
		fen.WriteString("/")
	}
	return fen.String()[:fen.Len()-1] + " w KQkq - 0 1"
}
