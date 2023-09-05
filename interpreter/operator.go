package interpreter

import (
	"fmt"
	"math"
	"strconv"

	"github.com/mariuskressin/dustmound/globals"
)

func ApplyOperator(o globals.Argument, a globals.Argument, b globals.Argument) (string, string) {
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
	case "&", "and":
		newType = "boolean"
		newVal = ToBool(a) && ToBool(b)
	case "!&", "nand":
		newType = "boolean"
		newVal = !(ToBool(a) && ToBool(b))
	case "|", "or":
		newType = "boolean"
		newVal = ToBool(a) || ToBool(b)
	case "||", "xor":
		newType = "boolean"
		if ToBool(a) && ToBool(b) {
			newVal = false
		} else {
			newVal = ToBool(a) || ToBool(b)
		}
	case "!|", "nor":
		newType = "boolean"
		newVal = !(ToBool(a) || ToBool(b))
	case ">", "gt":
		newType = "boolean"
		if (a.Type == "int" || a.Type == "float") && (b.Type == "int" || b.Type == "float") {
			aval, _ := strconv.ParseFloat(a.Value, 8)
			bval, _ := strconv.ParseFloat(b.Value, 8)
			newVal = aval > bval
		} else {
			fmt.Printf("\033[31mError:\033[32m Attempted comparison on non-number type: %s\033[0m\n", a.Type)
			panic("Bad types")
		}
	case "<", "lt":
		newType = "boolean"
		if (a.Type == "int" || a.Type == "float") && (b.Type == "int" || b.Type == "float") {
			aval, _ := strconv.ParseFloat(a.Value, 8)
			bval, _ := strconv.ParseFloat(b.Value, 8)
			newVal = aval < bval
		} else {
			fmt.Printf("\033[31mError:\033[32m Attempted comparison on non-number type: %s\033[0m\n", a.Type)
			panic("Bad types")
		}
	case ">=", "ge":
		newType = "boolean"
		if (a.Type == "int" || a.Type == "float") && (b.Type == "int" || b.Type == "float") {
			aval, _ := strconv.ParseFloat(a.Value, 8)
			bval, _ := strconv.ParseFloat(b.Value, 8)
			newVal = aval >= bval
		} else {
			fmt.Printf("\033[31mError:\033[32m Attempted comparison on non-number type: %s\033[0m\n", a.Type)
			panic("Bad types")
		}
	case "<=", "le":
		newType = "boolean"
		if (a.Type == "int" || a.Type == "float") && (b.Type == "int" || b.Type == "float") {
			aval, _ := strconv.ParseFloat(a.Value, 8)
			bval, _ := strconv.ParseFloat(b.Value, 8)
			newVal = aval <= bval
		} else {
			fmt.Printf("\033[31mError:\033[32m Attempted comparison on non-number type: %s\033[0m\n", a.Type)
			panic("Bad types")
		}
	case "=", "is":
		newType = "boolean"
		newVal = a.Value == b.Value
	case "==", "eq":
		newType = "boolean"
		newVal = (a.Value == b.Value) && (a.Type == b.Type)
	case "!=", "isnt":
		newType = "boolean"
		newVal = (a.Value != b.Value) || (a.Type != b.Type)
	case "!==", "neq":
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
	return newType, stringNewVal
}
