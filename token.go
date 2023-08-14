package main

import "regexp"

type Token struct {
  Type  string
	Value string
}

func BeginToken(char string, inside []string) (string, []string) {
	var token string
	var wordchar = regexp.MustCompile("\\w")
	var stringdelimiter = regexp.MustCompile("^[\"'`]$")
	if wordchar.MatchString(char) {
		token = char
		inside = append(inside, "keyword")
	} else if stringdelimiter.MatchString(char) {
		token = char
		inside = append(inside, char)
	}
	return token, inside
}
