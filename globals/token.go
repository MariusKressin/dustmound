package globals

type Token struct {
	Type      string
	Value     string
	ID        int
	BelongsTo int
}

func (t Token) Expr() Expression {
	return Expression{
		Type: t.Type,
		Name: t.Value,
		Args: make([]Argument, 0),
	}
}
