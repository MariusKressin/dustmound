package interpreter

import (
	"fmt"
	"math"
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
		if o.Type == "operator" && (o.Value == "!" || o.Value == "!!") {
			if i != len(evalArgs)-1 {
				newBool := "t"
				if ToBool(evalArgs[i+1]) && o.Value == "!" {
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
	i = 0
	for {
		o := evalArgs[i]
		if o.Type == "operator" {
			var a = globals.Argument{}
			if i != 0 {
				a = evalArgs[i-1]
			}
			var b = evalArgs[i+1]
			var newType = "string"
			var newVal any
			switch o.Value {
			case "+":
				if (a.Type == "float" || a.Type == "int") && (b.Type == "float" || b.Type == "int") {
					aval, _ := strconv.ParseFloat(a.Value, 8)
					bval, _ := strconv.ParseFloat(b.Value, 8)
					newVal = aval + bval
					if math.Round(aval+bval) == (aval + bval) {
						newVal = int(math.Round(aval + bval))
						newType = "int"
					} else {
						newType = "float"
					}
				} else if a.Type == "string" && b.Type == "string" {
					newVal = a.Value + b.Value
					newType = "string"
				} else if a.Type == "string" {
					fmt.Printf("\033[31mError:\033[32m Failed string concatenation. \033[0m\n")
					panic("Failed string concat")
				} else if a.Type == "int" || a.Type == "float" {
					fmt.Printf("\033[31mError:\033[32m Attempted to add %s to %s.\033[0m\n", b.Type, a.Type)
					panic("Bad types")
				} else {
					fmt.Printf("\033[31mError:\033[32m Attempted addition on unaddable type: %s\033[0m\n", a.Type)
					panic("Bad types")
				}
			case "-":
				if (a.Type == "float" || a.Type == "int") && (b.Type == "float" || b.Type == "int") {
					aval, _ := strconv.ParseFloat(a.Value, 8)
					bval, _ := strconv.ParseFloat(b.Value, 8)
					newVal = aval - bval
					if math.Round(aval-bval) == (aval - bval) {
						newVal = int(math.Round(aval - bval))
						newType = "int"
					} else {
						newType = "float"
					}
				} else if a.Type == "int" || a.Type == "float" {
					fmt.Printf("\033[31mError:\033[32m Attempted to subtract %s from %s.\033[0m\n", b.Type, a.Type)
					panic("Bad types")
				} else {
					fmt.Printf("\033[31mError:\033[32m Attempted subtraction on non-number type: %s\033[0m\n", a.Type)
					panic("Bad types")
				}
			case "*":
				if (a.Type == "float" || a.Type == "int") && (b.Type == "float" || b.Type == "int") {
					aval, _ := strconv.ParseFloat(a.Value, 8)
					bval, _ := strconv.ParseFloat(b.Value, 8)
					newVal = aval * bval
					if math.Round(aval*bval) == (aval * bval) {
						newVal = int(math.Round(aval * bval))
						newType = "int"
					} else {
						newType = "float"
					}
				} else if a.Type == "int" || a.Type == "float" {
					fmt.Printf("\033[31mError:\033[32m Attempted to multiply %s by %s.\033[0m\n", a.Type, b.Type)
					panic("Bad types")
				} else {
					fmt.Printf("\033[31mError:\033[32m Attempted multiplication on non-number type: %s\033[0m\n", a.Type)
					panic("Bad types")
				}
			case "/":
				if (a.Type == "float" || a.Type == "int") && (b.Type == "float" || b.Type == "int") {
					aval, _ := strconv.ParseFloat(a.Value, 8)
					bval, _ := strconv.ParseFloat(b.Value, 8)
					newVal = aval / bval
					if math.Round(aval/bval) == (aval / bval) {
						newVal = int(math.Round(aval / bval))
						newType = "int"
					} else {
						newType = "float"
					}
				} else if a.Type == "int" || a.Type == "float" {
					fmt.Printf("\033[31mError:\033[32m Attempted to divide %s by %s.\033[0m\n", a.Type, b.Type)
					panic("Bad types")
				} else {
					fmt.Printf("\033[31mError:\033[32m Attempted division on non-number type: %s\033[0m\n", a.Type)
					panic("Bad types")
				}
			case "%":
				if a.Type != b.Type {
					fmt.Printf("\033[31mError:\033[32m Mismatched types: %s and %s.\033[0m\n", a.Type, b.Type)
					panic("Mismatched types")
				} else if a.Type == "int" {
					ai, _ := strconv.Atoi(a.Value)
					bi, _ := strconv.Atoi(b.Value)
					newVal = ai % bi
				} else {
					fmt.Printf("\033[31mError:\033[32m Attempted modulo on non-int type: %s\033[0m\n", a.Type)
					panic("Bad types")
				}
				newType = a.Type
			case "^":
				if (a.Type == "float" || a.Type == "int") && (b.Type == "float" || b.Type == "int") {
					aval, _ := strconv.ParseFloat(a.Value, 8)
					bval, _ := strconv.ParseFloat(b.Value, 8)
					newVal = math.Pow(aval, bval)
					if math.Round(math.Pow(aval, bval)) == math.Pow(aval, bval) {
						newVal = int(math.Round(math.Pow(aval, bval)))
						newType = "int"
					} else {
						newType = "float"
					}
				} else if a.Type == "int" || a.Type == "float" {
					fmt.Printf("\033[31mError:\033[32m Attempted to take %s to the power of %s.\033[0m\n", a.Type, b.Type)
					panic("Bad types")
				} else {
					fmt.Printf("\033[31mError:\033[32m Attempted multiplication on non-number type: %s\033[0m\n", a.Type)
					panic("Bad types")
				}
			case "&":
				newType = "boolean"
				newVal = ToBool(a) && ToBool(b)
			case "!&":
				newType = "boolean"
				newVal = !(ToBool(a) && ToBool(b))
			case "|":
				newType = "boolean"
				newVal = ToBool(a) || ToBool(b)
			case "||":
				newType = "boolean"
				if ToBool(a) && ToBool(b) {
					newVal = false
				} else {
					newVal = ToBool(a) || ToBool(b)
				}
			case "!|":
				newType = "boolean"
				newVal = !(ToBool(a) || ToBool(b))
			case ">":
				newType = "boolean"
				if (a.Type == "int" || a.Type == "float") && (b.Type == "int" || b.Type == "float") {
					aval, _ := strconv.ParseFloat(a.Value, 8)
					bval, _ := strconv.ParseFloat(b.Value, 8)
					newVal = aval > bval
				} else {
					fmt.Printf("\033[31mError:\033[32m Attempted comparison on non-number type: %s\033[0m\n", a.Type)
					panic("Bad types")
				}
			case "<":
				newType = "boolean"
				if (a.Type == "int" || a.Type == "float") && (b.Type == "int" || b.Type == "float") {
					aval, _ := strconv.ParseFloat(a.Value, 8)
					bval, _ := strconv.ParseFloat(b.Value, 8)
					newVal = aval < bval
				} else {
					fmt.Printf("\033[31mError:\033[32m Attempted comparison on non-number type: %s\033[0m\n", a.Type)
					panic("Bad types")
				}
			case ">=":
				newType = "boolean"
				if (a.Type == "int" || a.Type == "float") && (b.Type == "int" || b.Type == "float") {
					aval, _ := strconv.ParseFloat(a.Value, 8)
					bval, _ := strconv.ParseFloat(b.Value, 8)
					newVal = aval >= bval
				} else {
					fmt.Printf("\033[31mError:\033[32m Attempted comparison on non-number type: %s\033[0m\n", a.Type)
					panic("Bad types")
				}
			case "<=":
				newType = "boolean"
				if (a.Type == "int" || a.Type == "float") && (b.Type == "int" || b.Type == "float") {
					aval, _ := strconv.ParseFloat(a.Value, 8)
					bval, _ := strconv.ParseFloat(b.Value, 8)
					newVal = aval <= bval
				} else {
					fmt.Printf("\033[31mError:\033[32m Attempted comparison on non-number type: %s\033[0m\n", a.Type)
					panic("Bad types")
				}
			case "=":
				newType = "boolean"
				newVal = a.Value == b.Value
			case "==":
				newType = "boolean"
				newVal = (a.Value == b.Value) && (a.Type == b.Type)
			case "!=":
				newType = "boolean"
				newVal = (a.Value != b.Value) || (a.Type != b.Type)
			case "!==":
				newType = "boolean"
				newVal = a.Value != b.Value
			}
			var stringNewVal = fmt.Sprintf("%s", newVal)
			switch newType {
			case "int":
				stringNewVal = fmt.Sprintf("%d", newVal)
			case "float":
				stringNewVal = fmt.Sprintf("%f", newVal)
			case "boolean":
				if newVal == true {
					stringNewVal = "t"
				} else {
					stringNewVal = "f"
				}
			}
			if i >= len(evalArgs)-1 {
				evalArgs = append(evalArgs[:i-1], globals.Argument{
					Type:  newType,
					Value: stringNewVal,
				})
			} else if i == 0 {
				evalArgs = append([]globals.Argument{{
					Type:  newType,
					Value: stringNewVal,
				}}, evalArgs[2:]...)
				i++
			} else {
				evalArgs = append(append(evalArgs[:i-1], globals.Argument{
					Type:  newType,
					Value: stringNewVal,
				}), evalArgs[i+2:]...)
			}
		} else {
			i++
		}
		if i >= len(evalArgs)-1 {
			break
		}
	}
	return evalArgs
}
