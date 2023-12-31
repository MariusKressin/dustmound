package lexer

import (
	"regexp"

	"github.com/mariuskressin/dustmound/globals"
)

// Variables for storing token-related information
var Inside = []globals.Token{} // Current scope
var Tokens []globals.Token     // List of tokens
var CurrentToken string        // Current token
var CurrentTokenID int = 1     // Next token ID

// Regexps
var WordChar = regexp.MustCompile("[a-zA-Z_0-9\\.]")                       // Regexp for word character
var StringDelimiter = regexp.MustCompile("^[\"'`]$")                       // Regexp for string delimiter
var WhiteSpace = regexp.MustCompile("[\\n\\t\\r ]")                        // Regexp for white space
var IntRegexp = regexp.MustCompile("^[0-9]+$")                             // Regexp for int
var FloatRegexp = regexp.MustCompile("^[0-9]+\\.[0-9]+$")                  // Regexp for float
var Operator = regexp.MustCompile("[\\+\\-\\=\\!\\<\\>\\*/\\&\\|%\\.\\^]") // Regexp for operator
