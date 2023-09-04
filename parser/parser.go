package parser

import (
	"strconv"

	"github.com/mariuskressin/dustmound/globals"
)

var MaxID = 0
var BlockLevel = 0
var mode = "normal"
var conditions = make([]globals.Condition, 0)
var currentCondition = globals.Condition{
	Type: "condition",
}

func ParseTokens(tokens []globals.Token) []*globals.Command {
	var LevelPtr = &BlockLevel
	// Loop over the tokens
	for _, t := range tokens {

		if mode == "condition" {
			if t.Type == "command" {
				currentCondition.Args = append(currentCondition.Args, globals.Argument{
					Type:  "eval",
					Value: strconv.Itoa(t.ID),
				})
			} else if t.Value != "do" {
				currentCondition.Args = append(currentCondition.Args, globals.Argument{
					Type:  t.Type,
					Value: t.Value,
				})
			}
		}
		if t.Type == "command" {
			if t.BelongsTo != 0 && mode != "condition" {
				for _, c := range Commands {
					if c.ArgsID == t.BelongsTo {
						c.Args = append(c.Args, globals.Argument{
							Type:  "eval",
							Value: strconv.Itoa(t.ID),
						})
					}
				}
			}
			Commands = append(Commands, &globals.Command{
				Name:      t.Value,
				ID:        t.ID,
				Args:      make([]globals.Argument, 0),
				ArgsID:    0,
				Condition: conditions,
				PassTo: globals.Location{
					ID:    t.BelongsTo,
					Index: 0,
				},
			})
		} else if t.Type == "keyword" {
			// Do some other stuff depending on the keyword.
			if t.Value == "if" {
				mode = "condition"
				currentCondition = globals.Condition{
					Type: "if",
					Args: make([]globals.Argument, 0),
				}
			} else if t.Value == "do" {
				if currentCondition.Type == "if" {
					mode = "normal"
					conditions = append(conditions, currentCondition)
					currentCondition = globals.Condition{}
					*LevelPtr++
				} else {
					panic("Unexpected \"do\".")
				}
			} else if t.Value == "end" {
				*LevelPtr--
				if BlockLevel < 0 {
					panic("Unexpected \"end\".")
				}
			}
		} else if t.Type == "list" {
			if t.BelongsTo != 0 {
				Arglists = append(Arglists, &globals.Arglist{
					Args: make([]string, 0),
					ID:   t.ID,
				})
				for _, c := range Commands {
					if c.ID == t.BelongsTo {
						c.ArgsID = t.ID
					}
				}
			}
		} else {
			if t.BelongsTo != 0 {
				for _, c := range Commands {
					if c.ArgsID == t.BelongsTo {
						c.Args = append(c.Args, globals.Argument{
							Type:  t.Type,
							Value: t.Value,
						})
					}
				}
			}
		}
	}

	return Commands
}
