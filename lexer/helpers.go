package lexer

import (
	"github.com/mariuskressin/dustmound/globals"
)

func BeginToken(char string, belongsTo int) {
	var belongsToID = or(belongsTo, Inside[len(Inside)-1].ID)
	if WordChar.MatchString(char) {
		CurrentToken = char
		Inside = append(Inside, Token{
			Type: "word",
			Value: "",
			ID: CurrentTokenID,
			BelongsTo: belongsToID,
		})
	} else if StringDelimiter.MatchString(char) {
		CurrentToken = char
		Inside = append(Inside, Token{
			Type: char,
			Value: "",
			ID: CurrentTokenID,
			BelongsTo: belongsToID,
		})
	} else if char == "(" {
		CurrentToken = ""
    Inside = append(Inside, Token{
      Type: "list",
      Value: "",
      ID: CurrentTokenID,
      BelongsTo: belongsToID,
    })
		CurrentTokenID ++
	} else if Operator.MatchString(char) {
		CurrentToken = char
		Inside = append(Inside, Token{
			Type: "operator",
			Value: char,
			ID: CurrentTokenID,
			BelongsTo: belongsToID,
		})
	}
}


func DetectWordType() string {
	var wordType = "identifier" // Default to identifier
	for i := range globals.Commands {   // Iterate over commands and check if the current token matches it.
		if CurrentToken == globals.Commands[i] {
			wordType = "command" // If the token matches a command, set the word type to command.
			break
		}
	}
	for i := range globals.Keywords { // Do the same thing with keywords
		if CurrentToken == globals.Keywords[i] {
			wordType = "keyword"
			break
		}
	}
	for i := range globals.Datatypes { // And datatypes
		if CurrentToken == globals.Datatypes[i] {
			wordType = "datatype"
			break
		}
	}
	for i := range globals.Operators { // And word-based operators
		if CurrentToken == globals.Operators[i] {
			wordType = "operator"
			break
		}
	}
	if CurrentToken == "t" || CurrentToken == "f" {
		wordType = "boolean"
	} else if CurrentToken == "nil" {
		wordType = "nil"
	} else if CurrentToken == "undef" {
		wordType = "undef"
	} else if IntRegexp.MatchString(CurrentToken) {
		wordType = "int"
	} else if FloatRegexp.MatchString(CurrentToken) {
		wordType = "float"
	}

	return wordType
}

func CreateToken(tokenType string) int {
	var value = "("
	var belongsTo = -1
	var id = CurrentTokenID
	for _, t := range Tokens {
		if t.BelongsTo == Inside[len(Inside)-1].ID {
			value += " " + t.Value
		}
	}
	value += " )"
	if tokenType != "list" {
		Inside = Inside[:len(Inside)-1]
		belongsTo = Inside[len(Inside)-1].ID
		value = CurrentToken
	} else {
		id = Inside[len(Inside)-1].ID
		belongsTo = Inside[len(Inside)-1].BelongsTo
	}
	Tokens = append(Tokens, Token{
		Type:      tokenType,
		Value:     value,
		ID:        id,
		BelongsTo: belongsTo,
	})
	if tokenType == "list" {
		Inside = Inside[:len(Inside)-1]
	} else {
		CurrentTokenID++
	}
	var nextTokenBelongsTo = -1
	if tokenType == "command" {
		nextTokenBelongsTo = CurrentTokenID - 1
	}
	return nextTokenBelongsTo
}

func or (first int, second int) int {
	if first == -1 {
		return second
	} else {
		return first
	}
}
