package main

import (
	"fmt"
	"regexp"
	"os"
)

func main() {
	if len(os.Args) < 2 {
    fmt.Println("Usage: go run main.go <file>")
    os.Exit(1)
  }

  file, err := os.ReadFile(os.Args[1])
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

	// Variable for storing code
	var code = string(file)

	// Variables for storing token-related information
	var inside = []Token{} // Current scope
	var tokens []Token // List of tokens
	var token string // Current token
	var currentTokenID int = 1 // Next token ID

	// Regexps
	var wordchar = regexp.MustCompile("[a-zA-Z_]") // Regexp for word character
	var stringdelimiter = regexp.MustCompile("^[\"'`]$") // Regexp for string delimiter

	inside = append(inside, Token{
		Type: "env",
		Value: "env",
		ID: 0,
		BelongsTo: -1,
	})

	for _, c := range code {
		// Begin tokens
		if inside[len(inside)-1].Type == "env" {
			token, inside = BeginToken(string(c), inside, currentTokenID)
		} else if inside[len(inside)-1].Type == "keyword" { // End/Continue keyword
			if wordchar.MatchString(string(c)) {
				token += string(c)
      } else {
				inside = inside[:len(inside)-1]
				tokens = append(tokens, Token{
					Type: "keyword",
					Value: token,
					ID: currentTokenID,
					BelongsTo: inside[len(inside)-1].ID,
				})
				currentTokenID ++
				token, inside = BeginToken(string(c), inside, currentTokenID)
      }
		} else if stringdelimiter.MatchString(inside[len(inside)-1].Type) { // End/Continue string
			if string(c) == inside[len(inside)-1].Type {
        inside = inside[:len(inside)-1];
				token += string(c)
				tokens = append(tokens, Token{
					Type: "string",
					Value: token,
					ID: currentTokenID,
					BelongsTo: inside[len(inside)-1].ID,
				})
				currentTokenID ++
				token = ""
			} else {
				token += string(c)
			}
		}
	}

	// Output tokens in a formatted list
	for _, t := range tokens {
    fmt.Printf("%d: %s %s -- %d\n", t.ID, t.Type, t.Value, t.BelongsTo)
  }
}
