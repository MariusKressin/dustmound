package main

import (
	"fmt"
	"regexp"
	"sort"
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

	// Variable for storing command names
	var commands = []string{
		"line",
		"semi_line",
	}

	// Variable for storing keyword names
	var keywords = []string{
		"set",
		"if",
		"then",
		"else",
		"end",
		"while",
		"until",
		"do",
		"for",
		"break",
    "continue",
		"com",
		"command",
		"out",
		"in",
		"to",
	}

	// Variables for storing token-related information
	var inside = []Token{} // Current scope
	var tokens []Token // List of tokens
	var token string // Current token
	var currentTokenID int = 1 // Next token ID

	// Regexps
	var wordchar = regexp.MustCompile("[a-zA-Z_0-9]") // Regexp for word character
	var stringdelimiter = regexp.MustCompile("^[\"'`]$") // Regexp for string delimiter
	var number = regexp.MustCompile("^[0-9]+$") // Regexp for number

	inside = append(inside, Token{
		Type: "env",
		Value: "env",
		ID: 0,
		BelongsTo: -1,
	})

	for _, c := range code {
		// Begin tokens
		if inside[len(inside)-1].Type == "env" {
			token, inside, currentTokenID = BeginToken(string(c), inside, currentTokenID, -1)
		} else if inside[len(inside)-1].Type == "word" { // End/Continue word
			if wordchar.MatchString(string(c)) {
				token += string(c)
      } else { // End word, and figure out what the type is
				var wordType = "identifier" // Default to identifier
				for i := range(commands) { // Iterate over commands and check if the current token matches it.
					if token == commands[i] {
            wordType = "command" // If the token matches a command, set the word type to command.
            break
          }
				}
				for i := range(keywords) { // Do the same thing with keywords
					if token == keywords[i] {
            wordType = "keyword"
            break
          }
				}
				if token == "t" || token == "f" {
					wordType = "boolean"
				} else if token == "nil" {
					wordType = "nil"
				} else if token == "undef" {
					wordType = "undef"
				} else if number.MatchString(token) {
					wordType = "number"
				}

				inside = inside[:len(inside)-1]

				tokens = append(tokens, Token{
					Type: wordType,
					Value: token,
					ID: currentTokenID,
					BelongsTo: inside[len(inside)-1].ID,
				})

				currentTokenID ++
				token, inside, currentTokenID = BeginToken(string(c), inside, currentTokenID, -1)
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
		} else if inside[len(inside)-1].Type == "list" {
			if string(c) == ")" {
        tokens = append(tokens, Token{
          Type: "list",
          Value: "",
          ID: inside[len(inside)-1].ID,
          BelongsTo: inside[len(inside)-2].ID,
        })
				inside = inside[:len(inside)-1]
			} else {
				token, inside, currentTokenID = BeginToken(string(c), inside, currentTokenID, -1)
			}
		}
	}

	// Sort the tokens by id
	sort.Slice(tokens, func(i, j int) bool {
    return tokens[i].ID < tokens[j].ID
	})

	// Output tokens in a formatted list
	for _, t := range tokens {
    fmt.Printf("\033[31m%-3d\033[97m>> %-3d \033[32m%-11s\033[37m%s\033[0m\n", t.ID, t.BelongsTo, t.Type, t.Value)
  }
}
