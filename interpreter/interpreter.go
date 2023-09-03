package interpreter

import (
	"github.com/mariuskressin/dustmound/globals"
)

func Interpret(commands []*globals.Command) {
	for _, c := range commands {
		if c.PassTo.ID == 0 {
			Eval(*c.Expr())
		}
	}
}
