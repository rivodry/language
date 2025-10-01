package parser

import (
	"lang/ast"
	"lang/lex"
)

func Parse(tokens []lex.Token) any {
	p := &Parser{tokens: tokens}
	node := p.parseExpr()
	return *node.(*ast.BinNode)
}
