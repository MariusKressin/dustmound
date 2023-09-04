package globals

type Condition struct {
	Type string
	Args []Argument
}

func CompressConditions(conds []Condition) []Argument {
	var args = make([]Argument, 0)
	for i, c := range conds {
		if i != 0 {
			args = append(args, Argument{
				Type:  "operator",
				Value: "&",
			})
		}
		args = append(args, c.Args...)
	}
	return args
}
