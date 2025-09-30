package ast

type BinNode struct {
	Op    string
	Left  any
	Right any
}
type NumNode struct {
	Value string
}
type IdentNode struct {
	Ident string
}
