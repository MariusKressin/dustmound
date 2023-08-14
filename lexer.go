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
	var inside = []string{ "env" } // Current scope
	var tokens []Token // List of tokens
	var token string // Current token

	// Regexps
	var wordchar = regexp.MustCompile("\\w") // Regexp for word character
	var stringdelimiter = regexp.MustCompile("^[\"'`]$") // Regexp for string delimiter

	for _, c := range code {
		// Begin tokens
		if inside[len(inside)-1] == "env" {
			token, inside = BeginToken(string(c), inside)
		} else if inside[len(inside)-1] == "keyword" { // End/Continue keyword
			if wordchar.MatchString(string(c)) {
				token += string(c)
      } else {
				inside = inside[:len(inside)-1]
				tokens = append(tokens, Token{Type: "keyword", Value: token})
				token, inside = BeginToken(string(c), inside)
      }
		} else if stringdelimiter.MatchString(inside[len(inside)-1]) { // End/Continue string
			if string(c) == inside[len(inside)-1] {
				token += string(c)
				tokens = append(tokens, Token{Type: "string", Value: token})
				token = ""
        inside = inside[:len(inside)-1];
			} else {
				token += string(c)
			}
		}
	}

	// Output tokens in a formatted list
	for i, t := range tokens {
    fmt.Printf("%d: %s %s\n", i, t.Type, t.Value)
  }
}
