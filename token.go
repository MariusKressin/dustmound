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
	var operator = regexp.MustCompile("[\\+\\-\\=\\!\\<\\>\\*/\\&\\|%\\.]") // Regexp for number
	var newID = id;
	var belongsToID = or(belongsTo, inside[len(inside)-1].ID)
	if wordchar.MatchString(char) {
		token = char
		inside = append(inside, Token{
			Type: "word",
			Value: "",
			ID: id,
			BelongsTo: belongsToID,
		})
	} else if stringdelimiter.MatchString(char) {
		token = char
		inside = append(inside, Token{
			Type: char,
			Value: "",
			ID: id,
			BelongsTo: belongsToID,
		})
	} else if char == "(" {
		token = ""
    inside = append(inside, Token{
      Type: "list",
      Value: "",
      ID: id,
      BelongsTo: belongsToID,
    })
		newID ++
	} else if operator.MatchString(char) {
		token = char
		inside = append(inside, Token{
			Type: "operator",
			Value: char,
			ID: id,
			BelongsTo: belongsToID,
		})
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
