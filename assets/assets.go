package assets

import (
	_ "embed"
)

//go:embed logo.txt
var logo string

func GetLogo() string {
	return logo
}
