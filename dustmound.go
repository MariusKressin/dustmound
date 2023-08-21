package main

import (
	"github.com/mariuskressin/dustmound/globals"
	"github.com/mariuskressin/dustmound/lexer"
	"github.com/mariuskressin/dustmound/parser"
)

func main() {
	globals.Init()
	tokens := lexer.Tokenize()
	parser.ParseTokens(tokens)
}
