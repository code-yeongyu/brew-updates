package history

import (
	"strings"
)

func CountBrewUpgrade(history string) int {
	if history == "" {
		return 0
	}
	return strings.Count(history, "brew upgrade")
}
