package lexer

import (
	"sort"

	"github.com/mariuskressin/dustmound/globals"
)

func Tokenize() []globals.Token {
	Inside = append(Inside, globals.Token{
		Type:      "env",
		Value:     "env",
		ID:        0,
		BelongsTo: -1,
	})

	for _, c := range globals.Code {
		// Begin tokens
		if Inside[len(Inside)-1].Type == "env" {
			BeginToken(string(c), -1)
		} else if Inside[len(Inside)-1].Type == "word" { // End/Continue word
			if WordChar.MatchString(string(c)) {
				CurrentToken += string(c)
			} else { // End word, and figure out what the type is
				var wordType = DetectWordType()
				var nextTokenBelongsTo = CreateToken(wordType)
				if string(c) == ")" {
					CreateToken("list")
				} else {
					BeginToken(string(c), nextTokenBelongsTo)
				}
			}
		} else if StringDelimiter.MatchString(Inside[len(Inside)-1].Type) { // End/Continue string
			if string(c) == Inside[len(Inside)-1].Type {
				CurrentToken += string(c)
				CreateToken("string")
			} else {
				CurrentToken += string(c)
			}
		} else if Inside[len(Inside)-1].Type == "list" {
			if string(c) == ")" {
				CreateToken("list")
			} else {
				BeginToken(string(c), -1)
			}
		} else if Inside[len(Inside)-1].Type == "operator" {
			if Operator.MatchString(string(c)) {
				CurrentToken += string(c)
			} else {
				CreateToken("operator")
				if string(c) == ")" {
					CreateToken("list")
				} else {
					BeginToken(string(c), -1)
				}
			}
		} else if Inside[len(Inside)-1].Type == "comment" {
			if string(c) == "\n" {
				Inside = Inside[:len(Inside)-1]
			}
		}
	}

	// Sort the tokens by id
	sort.Slice(Tokens, func(i, j int) bool {
		return Tokens[i].ID < Tokens[j].ID
	})

	return Tokens
}
