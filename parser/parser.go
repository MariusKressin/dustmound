package parser

import (
	"strconv"

	"github.com/mariuskressin/dustmound/globals"
)

var MaxID = 0

func ParseTokens(tokens []globals.Token) {
	// Loop over the tokens
	for _, t := range tokens {
		if t.Type == "command" {
			if t.BelongsTo != 0 {
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
				Name:   t.Value,
				ID:     t.ID,
				Args:   make([]globals.Argument, 0),
				ArgsID: 0,
				PassTo: globals.Location{
					ID:    t.BelongsTo,
					Index: 0,
				},
			})
		} else if t.Type == "keyword" {
			// Do some other stuff depending on the keyword.
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

	for _, c := range Commands {
		if c.PassTo.ID == 0 {
			Eval(*c.Expr())
		}
	}
}
