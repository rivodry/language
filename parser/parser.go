package parser

import (
	"lang/ast"
	"lang/lex"
)

func Parse(tokens []lex.Token) ast.BinNode {
	p := &Parser{tokens: tokens}
	node := p.parseExpr()
	return *node.(*ast.BinNode)
}
