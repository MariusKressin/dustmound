package parser

import (
	"fmt"
	"math"
	"strconv"

	"github.com/mariuskressin/dustmound/globals"
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
		for _, v := range Identifiers {
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

func CompressArgs(e globals.Expression) []globals.Argument {
	var evalArgs = make([]globals.Argument, 0)
	for _, a := range e.Args {
		var expression = a.Expr()
		var evaluated = globals.Argument{}
		if expression.Type == "eval" {
			for _, c := range Commands {
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
	i := 0
	for {
		o := evalArgs[i]
		if o.Type == "operator" {
			var a = evalArgs[i-1]
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
			}
			var stringNewVal = fmt.Sprintf("%s", newVal)
			switch newType {
			case "int":
				stringNewVal = fmt.Sprintf("%d", newVal)
			case "float":
				stringNewVal = fmt.Sprintf("%f", newVal)
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
