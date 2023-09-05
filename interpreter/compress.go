package interpreter

import (
	"strconv"

	"github.com/mariuskressin/dustmound/globals"
	"github.com/mariuskressin/dustmound/parser"
)

func CompressArgs(e globals.Expression, condition bool) []globals.Argument {
	var evalArgs = make([]globals.Argument, 0)
	var arguments = e.Args
	if condition {
		arguments = globals.CompressConditions(e.Condition)
	}
	if len(arguments) == 0 {
		return arguments
	}

	for _, a := range arguments {
		var expression = a.Expr()
		var evaluated = globals.Argument{}
		if expression.Type == "eval" {
			for _, c := range parser.Commands {
				if strconv.Itoa(c.ID) == expression.Name {
					evaluated = Eval(*c.Expr())
				}
			}
		}
		if evaluated == (globals.Argument{}) {
			evaluated = Eval(expression)
			if evaluated == (globals.Argument{}) {
				continue
			}
		}
		evalArgs = append(evalArgs, evaluated)
	}
	i := len(evalArgs) - 1
	for { // The "not" operator must be handled seperately, because it acts on the following value and can be chained (!!x)
		o := evalArgs[i]
		if o.Type == "operator" && (o.Value == "!" || o.Value == "not" || o.Value == "!!") {
			if i != len(evalArgs)-1 {
				newBool := "t"
				if ToBool(evalArgs[i+1]) && (o.Value == "!" || o.Value == "not") {
					newBool = "f"
				} else if !ToBool(evalArgs[i+1]) && o.Value == "!!" {
					newBool = "f"
				}
				evalArgs[i+1] = globals.Argument{
					Type:  "boolean",
					Value: newBool,
				}
				if i == 0 {
					evalArgs = evalArgs[1:]
				} else {
					evalArgs = append(evalArgs[:i-1], evalArgs[i+1:]...)
				}
			}
		}
		i--
		if i < 0 {
			break
		}
	}
	for _, op := range OrderOfOps {
		i = 0
		for {
			o := evalArgs[i]
			if o.Type == "operator" && o.Value == op {
				var a = globals.Argument{}
				if i != 0 {
					a = evalArgs[i-1]
				}
				var b = evalArgs[i+1]
				var newType, newVal = ApplyOperator(o, a, b)
				if i >= len(evalArgs)-1 {
					evalArgs = append(evalArgs[:i-1], globals.Argument{
						Type:  newType,
						Value: newVal,
					})
				} else if i == 0 {
					evalArgs = append([]globals.Argument{{
						Type:  newType,
						Value: newVal,
					}}, evalArgs[2:]...)
					i++
				} else {
					evalArgs = append(append(evalArgs[:i-1], globals.Argument{
						Type:  newType,
						Value: newVal,
					}), evalArgs[i+2:]...)
				}
			} else {
				i++
			}
			if i >= len(evalArgs)-1 {
				break
			}
		}
	}
	return evalArgs
}
