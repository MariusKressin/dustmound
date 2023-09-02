package globals

type Argument struct {
	Type  string
	Value string
}

func (a Argument) Expr() Expression {
	return Expression{
		Type: a.Type,
		Name: a.Value,
	}
}
