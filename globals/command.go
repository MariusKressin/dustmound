package globals

type Command struct {
	Name      string
	Args      []Argument
	Condition []Condition
	ArgsID    int
	ID        int
	PassTo    Location
}

func (c *Command) Expr() *Expression {
	return &Expression{
		Type:      "command",
		Name:      c.Name,
		Args:      c.Args,
		Condition: c.Condition,
	}
}
