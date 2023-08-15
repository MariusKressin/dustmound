package main

import "regexp"

type Token struct {
  Type      string
	Value     string
	ID        int
	BelongsTo int
}

func BeginToken(char string, inside []Token, id int, belongsTo int) (string, []Token, int) {
	var token string
	var wordchar = regexp.MustCompile("\\w")
	var stringdelimiter = regexp.MustCompile("^[\"'`]$")
	var newID = id;
	if wordchar.MatchString(char) {
		token = char
		inside = append(inside, Token{
			Type: "word",
			Value: "",
			ID: id,
			BelongsTo: or(belongsTo, inside[len(inside) - 1].ID),
		})
	} else if stringdelimiter.MatchString(char) {
		token = char
		inside = append(inside, Token{
			Type: char,
			Value: "",
			ID: id,
			BelongsTo: or(belongsTo, inside[len(inside) - 1].ID),
		})
	} else if char == "(" {
		token = ""
    inside = append(inside, Token{
      Type: "list",
      Value: "",
      ID: id,
      BelongsTo: or(belongsTo, inside[len(inside) - 1].ID),
    })
		newID ++
	}
	return token, inside, newID
}

func or (first int, second int) int {
	if first == -1 {
		return second
	} else {
		return first
	}
}
