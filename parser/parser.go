package parser

import (
	"fmt"

	"github.com/mariuskressin/dustmound/globals"
)

func ParseTokens (tokens []globals.Token) {
	// Output tokens in a formatted list
	for _, t := range tokens {
		fmt.Printf("\033[33m%-3d\033[0m: %-3d \033[32m%-11s\033[37m%s\033[0m\n", t.ID, t.BelongsTo, t.Type, t.Value)
	}
}
