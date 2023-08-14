package main

import "regexp"

type Token struct {
  Type      string
	Value     string
	ID        int
	BelongsTo int
}

func BeginToken(char string, inside []Token, id int) (string, []Token) {
	var token string
	var wordchar = regexp.MustCompile("\\w")
	var stringdelimiter = regexp.MustCompile("^[\"'`]$")
	if wordchar.MatchString(char) {
		token = char
		inside = append(inside, Token{
			Type: "word",
			Value: "",
			ID: id,
			BelongsTo: inside[len(inside) - 1].ID,
		})
	} else if stringdelimiter.MatchString(char) {
		token = char
		inside = append(inside, Token{
			Type: char,
			Value: "",
			ID: id,
			BelongsTo: inside[len(inside) - 1].ID,
		})
	}
	return token, inside
}
