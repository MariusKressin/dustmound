package parser

import (
	"fmt"

	"github.com/mariuskressin/dustmound/globals"
)

func Eval(e globals.Expression) any {
	if e.Type == "command" {
		if e.Name == "line" {
			var stringArgs = ""
			for _, a := range e.Args {
				var evaluated = Eval(a.Expr())
				if evaluated == nil {
					continue
				}
				stringArgs += fmt.Sprintf("%s ", evaluated)
			}
			fmt.Printf("%v\n", stringArgs)
			return fmt.Sprintf("%v", stringArgs)
		} else if e.Name == "semi_line" {
			var stringArgs = ""
			for _, a := range e.Args {
				var evaluated = Eval(a.Expr())
				if evaluated == nil {
					continue
				}
				stringArgs += fmt.Sprintf("%s ", evaluated)
			}
			fmt.Printf("%v", stringArgs)
			return fmt.Sprintf("%v", stringArgs)
		}
	} else if e.Type == "string" || e.Type == "int" || e.Type == "float" {
		return e.Name
	} else if e.Type == "identifier" {
		for _, v := range Identifiers {
			if v.Name == e.Name {
				return v.Value
			}
		}
		fmt.Printf("\033[31mError:\033[32m \"%s\" is undefined! \033[0m\n", e.Name)
		panic("Undeclared identifier")
	}
	return nil
}

func ParseList(id int, tokens []globals.Token) []globals.Argument {
	var list = make([]globals.Argument, 0)
	for _, t := range tokens {
		if t.BelongsTo == id {
			list = append(list, globals.Argument{
				Type:  t.Type,
				Value: t.Value,
			})
		}
	}
	return list
}
