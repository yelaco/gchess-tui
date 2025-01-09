package utils

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GeneratePlayerId() string {
	return uuid.NewString()
}

func GenerateBlankLine(count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat("\n", count)
}

func BoardToFen(board [][]string) string {
	var fen strings.Builder

	for _, row := range board {
		emptyCount := 0
		for _, square := range row {
			if square == "." || square == "" { // Treat "." or "" as empty square
				emptyCount++
			} else {
				if emptyCount > 0 {
					fen.WriteString(strconv.Itoa(emptyCount))
					emptyCount = 0
				}
				switch square {
				case "♖":
					fen.WriteString("r")
				case "♘":
					fen.WriteString("n")
				case "♗":
					fen.WriteString("b")
				case "♕":
					fen.WriteString("q")
				case "♔":
					fen.WriteString("k")
				case "♙":
					fen.WriteString("p")
				case "♜":
					fen.WriteString("R")
				case "♞":
					fen.WriteString("N")
				case "♝":
					fen.WriteString("B")
				case "♛":
					fen.WriteString("Q")
				case "♚":
					fen.WriteString("K")
				case "♟":
					fen.WriteString("P")
				default:
					panic("Invalid piece symbol: " + square)
				}
			}
		}
		if emptyCount > 0 {
			fen.WriteString(strconv.Itoa(emptyCount))
		}
		fen.WriteString("/") // Add row separator
	}

	fenStr := fen.String()
	return fenStr[:len(fenStr)-1] // Remove trailing "/"
}
