package interpreter

import (
	"fmt"

	"github.com/mariuskressin/dustmound/globals"
	"github.com/mariuskressin/dustmound/parser"
)

func Eval(e globals.Expression) globals.Argument {
	if e.Type == "command" {
		var args = CompressArgs(e)
		if e.Name == "line" {
			for i, a := range args {
				if i != 0 {
					fmt.Print(" ")
				}
				fmt.Print(a.Value)
			}
			fmt.Print("\n")
		} else if e.Name == "semi_line" {
			for i, a := range args {
				if i != 0 {
					fmt.Print(" ")
				}
				fmt.Print(a.Value)
			}
		}
	} else if e.Type == "identifier" {
		for _, v := range parser.Identifiers {
			if v.Name == e.Name {
				return globals.Argument{
					Type:  v.Type,
					Value: v.Value,
				}
			}
		}
		fmt.Printf("\033[31mError:\033[32m \"%s\" is undefined! \033[0m\n", e.Name)
		panic("Undeclared identifier")
	}
	return globals.Argument{
		Type:  e.Type,
		Value: e.Name,
	}
}
