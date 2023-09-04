package globals

type Expression struct {
	Type      string
	Name      string
	Args      []Argument
	Condition []Condition
}
