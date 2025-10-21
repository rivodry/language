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
type UnaryNode struct {
	Op    string
	Value any
}
type ExpressionStatement struct {
	Expr any
}
type VarNode struct {
	Name string
	Init any
}
type SetNode struct {
	Name string
	Set  any
}
type IfNode struct {
	If   any
	Then []any
	Else []any
}
