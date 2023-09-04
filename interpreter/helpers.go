package interpreter

import (
	"strconv"

	"github.com/mariuskressin/dustmound/globals"
)

func ToBool(a globals.Argument) bool {
	if a.Type == "int" || a.Type == "float" {
		float, _ := strconv.ParseFloat(a.Value, 8)
		if float == 0.0 {
			return false
		}
	} else if a.Type == "string" {
		if a.Value == "" {
			return false
		}
	} else if a.Type == "boolean" {
		if a.Value == "f" {
			return false
		}
	} else if a.Type == "nil" || a.Type == "undef" {
		return false
	}
	return true
}
