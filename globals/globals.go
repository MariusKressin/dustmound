package globals

import (
	"fmt"
	"os"
)

// Variable for storing command names
var Commands = []string{
	"line",
	"semi_line",
}

// Variable for storing keyword names
var Keywords = []string{
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

// Variable for storing data types
var Datatypes = []string{
	"int",
	"float",
	"string",
	"bool",
	"Command",
	"list",
	"block",
}

// Variable for storing operators
var Operators = []string{
	"+",  // Plus
	"-",  // Minus
	"*",  // Times
	"/",  // Divided by
	"^",  // To the power of
	"%",  // Modulus
	":",  // Child block a:b is the equivalent of a.b in Go.
	">",  // Greater than
	"<",  // Less than
	">=", // Greater than or equal to
	"<=", // Less than or equal to
	"=",  // Equal to
	"!=", // Not equal to
	"!",  // Not
	"||", // XOR
	"|",  // OR
	"!|", // NOR
	"&",  // AND
	"!&", // NAND
	"**", // Wild card e.g. a ** b returns true if a contains b
	"is",
	"not",
	"and",
	"or",
	"any",
	"ge",
	"le",
	"eq",
	"lt",
	"gt",
}

var Code string

func Init() {
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
	Code = string(file)
}
